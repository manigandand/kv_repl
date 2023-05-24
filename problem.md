## Exercise: K/V REPL with nested transactions

In this exercise, we ask you to write a command-line REPL (read-eval-print loop) that drives a simple in-memory key/value storage system. This system should also allow for nested transactions. A transaction can then be committed or aborted.

We realize that your time is valuable. Please try to limit this to 1-2 hours. If you can't complete the exercise in this time, please share what you have as a basis for a discussion.

Use whatever programming language, tools, and development environment you're most comfortable with. Be prepared to talk about your submission; you and one of our engineers will be walking through it.

### Submitting Code

Please do not host your code in a public GitHub repo. We want this exercise to be usable by
multiple people and don't want an answer easily searchable.

If your source is a single file, you can just put it inline in an email message. Otherwise, please
attach an archive (.zip or .tar.gz) of the source.

### Example Run

$ my-program

```shell
> WRITE a hello
> READ a
hello

> START
> WRITE a hello-again
> READ a
hello-again

> START
> DELETE a
> READ a
Key not found: a

> COMMIT
> READ a
Key not found: a

> WRITE a once-more
> READ a
once-more

> ABORT
> READ a
hello

> QUIT
Exiting...

```

### Commands

● **READ** <key>; Reads and prints, to stdout, the val associated with key. If the value is not
present an error is printed to stderr.

● **WRITE** <key>; <val>; Stores val in key.

● **DELETE** <key>; Removes all key from store. Future READ commands on that key will
return an error.

● **START** Start a transaction.

● **COMMIT** Commit a transaction. All actions in the current transaction are committed to
the parent transaction or the root store. If there is no current transaction an error is
output to stderr.

● **ABORT** Abort a transaction. All actions in the current transaction are discarded.

● **QUIT** Exit the REPL cleanly. A message to stderr may be output.

### Other details

● For simplicity, all keys and values are simple ASCII strings delimited by whitespace. No
quoting is needed.

● All errors are output to stderr.

● Commands are case-insensitive.

● As this is a simple command line program with no networking, there is only one &quot;client&quot;
at a time. There is no need for locking or multiple threads.
