# Terraform Provider for Uptrends

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

### Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.12.x
- [Go](https://golang.org/doc/install) 1.14 (to build the provider plugin)

## Using the provider

After building the provider, follow the instructions to [install it as a plugin](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin). After placing the provider your plugins directory, run `terraform init` to initialize it.

#### Examples

To create monitor in Uptrends, create a terraform resource like the following, and run `terraform apply`.

```hcl
# main.tf
resource "uptrends_monitor" "foo" {
    monitor_type = "Http"
    name         = "foo-http-monitor"
    url          = "https://example.com"
}
```

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.14+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

#### Building
Clone repository to: `$GOPATH/src/github.com/craigsands/terraform-provider-uptrends`

```sh
$ mkdir -p $GOPATH/src/github.com/craigsands;
$ cd $GOPATH/src/github.com/craigsands
$ git clone git@github.com:craigsands/terraform-provider-uptrends.git
```

Enter the provider directory and build the provider. To compile the provider, run `go build`. This will build the provider and put the provider binary in the current directory.

```sh
$ cd $GOPATH/src/github.com/craigsands/terraform-provider-uptrends
$ go build
```

#### Testing
In order to test the provider, run `go test`. This will run the unit test suite.

```sh
$ go test -v ./...
```

In order to run the full suite of Acceptance tests, set the environment variable `TF_ACC=true`.

*Note:* Acceptance tests *create real resources*, and often cost money to run. The environment variables `UPTRENDS_USERNAME` and `UPTRENDS_PASSWORD` must also be set with your associated keys for acceptance tests to work properly.
