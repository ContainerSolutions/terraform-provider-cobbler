# terraform-provider-cobbler
Terraform provider for Cobbler

Check the `config.tf.example` file to see how to configure this.

Create a `config.tf` file and then run
`$ terraform apply`


### Features

- [x] Create System
- [x] Destroy System
- [ ] Update System
- [ ] Add support for creating distros. It would be nice if we could just simply pass along the URL to
      the ISO file and the provider took care of the rest (downloading, mounting, importing, unmounting,
      deleting temp files...).
