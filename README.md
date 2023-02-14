# go-to-dart

Go-to-dart helps you convert Go structs to Dart classes that can be used with [json_serializable](https://pub.dev/packages/json_serializable).

## Features

- Supports only structs in the same package (no generics or embedded structs yet)
- Supports primitives, slices, maps, and pointers
- Support some other arbitrary types such as `time.Time` and `mo.Option` (easy to extend!)

Need something more? Please open an issue or even better, a PR!

## Installation

```bash
go install github.com/11wizards/go-to-dart
```

The above command will install go-to-dart in your `$GOPATH/bin` directory. Make sure that directory is in your `$PATH`.

## Usage

```bash
go-to-dart -i ./examples/user -o ./examples/user
```

## Example

Running the command above would take the package `./examples/user` below and generate a file `./examples/user/user.dart`.

```go
package user

import (
	"time"
)

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Options   map[string]string
	Tags      []string
}
```

Contents of `./examples/user/user.dart`:
```dart
import 'package:json_annotation/json_annotation.dart';

part 'user.g.dart';

@JsonSerializable()
class User {
  @JsonKey(name: 'ID') final int id;
  @JsonKey(name: 'Name') final String name;
  @JsonKey(name: 'Email') final String email;
  @JsonKey(name: 'Password') final String password;
  @JsonKey(name: 'CreatedAt') final DateTime createdAt;
  @JsonKey(name: 'UpdatedAt') final DateTime updatedAt;
  @JsonKey(name: 'DeletedAt') final DateTime? deletedAt;
  @JsonKey(name: 'Options') final Map<String, String> options;
  @JsonKey(name: 'Tags') final List<String> tags;

  User({
    required this.id,
    required this.name,
    required this.email,
    required this.password,
    required this.createdAt,
    required this.updatedAt,
    this.deletedAt,
    required this.options,
    required this.tags,
  });

  Map<String, dynamic> toJson() => _$UserToJson(this);

  factory User.fromJson(Map<String, dynamic> json) => _$UserFromJson(json);
}


```