qif-processor
=============

qif-processor is a command-line bulk processor for .qif (Quicken Interchange Format) files.

.qif is one of the most common data formats for personal finance software and most banks provide a .qif download.  Even though .qif is considered a legacy format, it is still widely used.  

When importing data into a personal finance software package there is often the need to update or clean the data in bulk.  qif-processor is an easy to use command-line bulk processor.

Working with .qif files
-----------------------
qif-processor executes a command on an input .qif file and produces an out .qif file.  To see a set of commands run qif-processor without any arguments.
```
qp
```
By default qif-processor reads Stdin for input .qif and writes out .qif Stdout.  You can use the --input and --output global flags to set input and output files.
```
qp --input inputfile.qif --output outputfile.qif <command> <options>
```
In Unix operating systems you can pipe multiple qif-processor runs together.
```
cat input.qif | qp <command> <options> | qp <command> <options> > output.qif
```

Deleting Transactions
---------------------
qif-processor allows you to delete a set of transactions from a qif file.

For example, to delete transactions from a .qif file that contain the string 'Some Corporation' in the payee field:
```
qp -I input.qif -O output.qif delete -P 'Some Corporation'
```
The -P flag tags a regular expression.  You can for example remove all transactions that starts with ATM:
```
qp -I input.qif -O output.qif delete -P '^ATM.*'
```

Setting Fields
--------------
You can set fields on transactions that match a regular expression.

For example to set the category of all transaction that match Woolworths to Everyday Expenses:Groceries you can use the following:
```
qp -I input.qif -O output.qif set-value -P 'Woolworths' -l 'Everyday Expenses:Groceries'
```

Developing and Building
-----------------------
qif-processor is written in Go.  To compile and build qif-processor get Go from https://golang.org/
```
git clone git@github.com:leonm/qif-processor.git
cd qif-processor
export GOPATH=`pwd`
go test ./...
go install ./...
./bin/qp
```
