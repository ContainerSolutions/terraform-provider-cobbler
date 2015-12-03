package main

import (
	"errors"
	"io/ioutil"

	cobbler "github.com/ContainerSolutions/cobblerclient"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCobblerSnippet() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        resourceCobblerSnippetCreate,
		Read:          resourceCobblerSnippetRead,
		Update:        resourceCobblerSnippetUpdate,
		Delete:        resourceCobblerSnippetDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name that the snippet file will be stored with in Cobbler",
			},
			"path": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Local path to the snippet file that will be created in Cobbler",
			},
			"version": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Bumpable version for when updates to the snippet file are needed",
			},
		},
	}
}

// This function is in charge of creating and/or updating a snippet file in Cobbler.
// The `cobbler.Snippet` struct consists of two fields: the `Name` that the file will
// have on the Cobbler server and the `Body` which must be the contents of the snippet
// file stored somewhere in the file directory from where terraform is ran.
// That means that we need to open given file from within this function. If the file
// is not found the function early returns a File Not Found kind of error.
func resourceCobblerSnippetCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cobbler.Client)
	ok, err := client.Login()
	if err != nil {
		return err
	}

	b, err := ioutil.ReadFile(d.Get("path").(string))
	if err != nil {
		return err
	}

	s := cobbler.Snippet{
		Name: d.Get("name").(string),
		Body: string(b),
	}
	ok, err = client.CreateSnippet(&s)
	if err != nil {
		return err
	}

	if ok {
		d.SetId(s.Name)
		return nil
	} else {
		return errors.New("could not create/update snippet file")
	}
}

func resourceCobblerSnippetRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceCobblerSnippetUpdate(d *schema.ResourceData, meta interface{}) error {
	if !d.HasChange("version") {
		return nil
	}

	return resourceCobblerSnippetCreate(d, meta)
}

func resourceCobblerSnippetDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
