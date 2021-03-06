/*
Copyright 2015 Container Solutions

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
			"version": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Bumpable version for when updates to the kickstart file are needed",
			},
		},
	}
}

// This function is in charge of creating and/or updating a kickstart file in Cobbler.
// The `cobbler.KickstartFile` struct consists of two fields: the `Name` that the file will
// have on the Cobbler server and the `Body` which must be the contents of the kickstart
// file stored somewhere in the file directory from where terraform is ran.
// That means that we need to open given file from within this function. If the file
// is not found the function early returns a File Not Found kind of error.
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
		d.SetId(ksf.Name)
		return nil
	} else {
		return errors.New("could not create/update kickstart file")
	}
}

func resourceCobblerKickstartFileRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceCobblerKickstartFileUpdate(d *schema.ResourceData, meta interface{}) error {
	if !d.HasChange("version") {
		return nil
	}

	return resourceCobblerKickstartFileCreate(d, meta)
}

func resourceCobblerKickstartFileDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
