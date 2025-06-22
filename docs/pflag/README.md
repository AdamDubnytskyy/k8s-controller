## [pflag](https://github.com/spf13/pflag)
pflag is a drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags.

### Usage
```sh
import flag "github.com/spf13/pflag"
```

There is one exception to this: if you directly instantiate the Flag struct there is one more field "Shorthand" that you will need to set. Most code never instantiates this struct directly, and instead uses functions such as String(), BoolVar(), and Var(), and is therefore unaffected.

Define flags using flag.String(), Bool(), Int(), etc.

This declares an integer flag, -flagname, stored in the pointer ip, with type *int.
```sh
var ip *int = flag.Int("flagname", 1234, "help message for flagname")
```

If you like, you can bind the flag to a variable using the Var() functions.
```sh
var flagvar int
func init() {
    flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
}
```

Or you can create custom flags that satisfy the Value interface (with pointer receivers) and couple them to flag parsing by
```sh
flag.Var(&flagVal, "name", "help message for flagname")
```

For such flags, the default value is just the initial value of the variable.

After all flags are defined, call
```sh
flag.Parse()
```
to parse the command line into the defined flags.

Flags may then be used directly. If you're using the flags themselves, they are all pointers; if you bind to variables, they're values.
```sh
fmt.Println("ip has value ", *ip)
fmt.Println("flagvar has value ", flagvar)
```

There are helper functions available to get the value stored in a Flag if you have a FlagSet but find it difficult to keep up with all of the pointers in your code. If you have a pflag.FlagSet with a flag called 'flagname' of type int you can use GetInt() to get the int value. But notice that 'flagname' must exist and it must be an int. GetString("flagname") will fail.
```sh
i, err := flagset.GetInt("flagname")
```

After parsing, the arguments after the flag are available as the slice flag.Args() or individually as flag.Arg(i). The arguments are indexed from 0 through flag.NArg()-1.

The pflag package also defines some new functions that are not in flag, that give one-letter shorthands for flags. You can use these by appending 'P' to the name of any function that defines a flag.
```sh
var ip = flag.IntP("flagname", "f", 1234, "help message")
var flagvar bool
func init() {
	flag.BoolVarP(&flagvar, "boolname", "b", true, "help message")
}
flag.VarP(&flagVal, "varname", "v", "help message")
```

Shorthand letters can be used with single dashes on the command line. Boolean shorthand flags can be combined with other shorthand flags.

The default set of command-line flags is controlled by top-level functions. The FlagSet type allows one to define independent sets of flags, such as to implement subcommands in a command-line interface. The methods of FlagSet are analogous to the top-level functions for the command-line flag set.

### Command line flag syntax
```sh
--flag    // boolean flags, or flags with no option default values
--flag x  // only on flags without a default value
--flag=x
```

Flag parsing stops after the terminator "--". Unlike the flag package, flags can be interspersed with arguments anywhere on the command line before this terminator.

Integer flags accept 1234, 0664, 0x1234 and may be negative. Boolean flags (in their long form) accept 1, 0, t, f, true, false, TRUE, FALSE, True, False. Duration flags accept any input valid for time.ParseDuration.


### Dynamically set log-level, cluster-name flags
```sh
./controller --log-level trace --cluster-name prod
```