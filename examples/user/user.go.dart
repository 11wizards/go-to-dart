import 'package:json_annotation/json_annotation.dart';

part 'user.go.g.dart';

@JsonSerializable()
class User {
	@JsonKey(name: "ID")final int id;
	@JsonKey(name: "Name")final String name;
	@JsonKey(name: "Email")final String email;
	@JsonKey(name: "Password")final String password;
	@JsonKey(name: "CreatedAt")final DateTime createdAt;
	@JsonKey(name: "UpdatedAt")final DateTime updatedAt;
	@JsonKey(name: "DeletedAt")final DateTime? deletedAt;
	@JsonKey(name: "Options", defaultValue: <String, String>{})final Map<String, String> options;
	@JsonKey(name: "Tags", defaultValue: <List<String>>[])final List<String> tags;
	
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

