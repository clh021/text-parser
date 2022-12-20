package main

type config struct {
	Format  string `mapstructure:"format"`
	Command string `mapstructure:"command"`
}
