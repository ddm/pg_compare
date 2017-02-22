# pg_compare

PGCompare compares 2 postgres databases with eachother.

## Comparison options

### Rows

Does a count of the amount of rows in a table and compares them with the 2 given databases.

### ForeignKeys

Checks that all foreign keys are present in both databases.

### Columns

Checks that the table has the same columns with the same settings.

### Views

Checks that all the views are present in both databases.

### Indexes

Checks that all the indexes are present in both databases.

## Installation

```
go install github.com/jelmersnoeck/pg_compare/cmd/pg_compare
```
