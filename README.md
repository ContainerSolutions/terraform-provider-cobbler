# terraform-provider-cobbler
A Terraform provider for Cobbler

The terraform-provider-cobbler allows for communication between Terraform and Cobbler. It was created to 
aid in the provisioning of machines where the available infrastructure does not provide dhcp or pxe-like 
services itself. An obvious use case is bare-metal server racks. A Cobbler instance is connected to the 
infrastructure and the Terraform provider allows for management of systems in Cobbler. Cobbler in turn 
provides the services needed for the provisioning of the machines: ip address, hostname, pxe boot image, 
kickstart file.

Currently the provider only allows for systems to be created and deleted. 
DHCP settings, software distributions, kickstart files and profiles should all be provided. 
Some of these are on the roadmap to be managed through Terraform.

This provider uses the [Cobblerclient](https://github.com/ContainerSolutions/cobblerclient "Cobblerclient") to talk to Cobbler.

### Configuration
Check the `config.tf.example` file to see how to configure this.

### Usage
Create a `config.tf` file and then run
`$ terraform apply`


### Features / Roadmap

- [x] Create System
- [x] Destroy System
- [ ] Update System
- [ ] Support for creating distros. It would be nice if we could just simply pass along the URL to
      the ISO file and the provider took care of the rest (downloading, mounting, importing, unmounting,
      deleting temp files...).
- [ ] Support for managing profiles
 
