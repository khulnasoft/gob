package gob

import "github.com/khulnasoft/gfi"

// FieldDescriber a struct implementing this interface will have DescribeFields called when the config is summarized
type FieldDescriber interface {
	DescribeFields(descriptions FieldDescriptionSet)
}

// FieldDescriptionSet FieldDescriber.DescribeFields will be called with this interface to add field descriptions
type FieldDescriptionSet = gfi.FieldDescriptionSet

// FlagAdder interface can be implemented by structs in order to add flags when AddFlags is called
type FlagAdder interface {
	AddFlags(flags FlagSet)
}

// FlagSet is a facade of pflag.FlagSet, restricting the types of calls to what gfi needs
type FlagSet = gfi.FlagSet

// PostLoader is the interface used to do any sort of processing after the entire struct has been populated
// from the configuration files and environment variables
type PostLoader = gfi.PostLoader
