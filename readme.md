### Family Tree CLI

> Prerequites
>
> 1) Postgres

- Once the postgres is running, update the config.yaml

## Building the Project

- Run go build,  this will generate an executable( family-tree ) in root folder

## Commands

> All command require --user="`<youruser>`" flag to be attached. This is so other users can create/add their family

Below command create a new user if it doesnt exist

```
./family-tree login --user="<user>" 
```

Family member name is not associated with above user. Both are independent. 

```
/family-tree login --user=anchit add person --person="<family-member-name>"--sex=male
```

User will define all the relationships within family. Current implementation has hardcoded mapping to to "father", "son", "daughter", "mother" which is used to find reverse relations

```
./family-tree login --user=anchit add relationship --relationship=son   
```

Count Query

```
 ./family-tree login --user=anchit count son of "<family-member-name>"  
```
