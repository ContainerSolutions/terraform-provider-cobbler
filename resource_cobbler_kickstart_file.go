package main

import (
	"errors"
	"io/ioutil"

	cobbler "github.com/ContainerSolutions/cobblerclient"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCobblerKickstartFile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        resourceCobblerKickstartFileCreate,
		Read:          resourceCobblerKickstartFileRead,
		Update:        resourceCobblerKickstartFileUpdate,
		Delete:        resourceCobblerKickstartFileDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name that the kickstart file will be stored with in Cobbler",
			},
			"path": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Local path to the kickstart file that will be created in Cobbler",
			},
		},
	}
}

func resourceCobblerKickstartFileCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cobbler.Client)
	ok, err := client.Login()
	if err != nil {
		return err
	}

	b, err := ioutil.ReadFile(d.Get("path").(string))
	if err != nil {
		return err
	}

	ksf := cobbler.KickstartFile{
		Name: d.Get("name").(string),
		Body: string(b),
	}
	ok, err = client.CreateKickstartFile(&ksf)
	if err != nil {
		return err
	}

	if ok {
		return nil
	} else {
		return errors.New("could not create kickstart file")
	}
}

func resourceCobblerKickstartFileRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceCobblerKickstartFileUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceCobblerKickstartFileDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
