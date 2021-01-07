# Tensorflow v2.4.0 Go Bindings

/!\ Not tested /!\

This is a fork of tensorflow's repository (at version 2.4.0) with pre-generated protobuf definitions, following instructions at https://github.com/tensorflow/tensorflow/issues/41808#issuecomment-683362089.

Inspired from https://github.com/zia-ai/tensorflow-go.

## Usage

Add this at the end of your go.mod file:

```replace github.com/tensorflow/tensorflow => github.com/onfocusio/tensorflow-go.v2.4.0 latest```

Then run `go mod tidy`.