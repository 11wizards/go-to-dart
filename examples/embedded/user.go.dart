// ignore_for_file: always_use_package_imports
import 'package:equatable/equatable.dart';
import 'package:json_annotation/json_annotation.dart';

part 'user.go.g.dart';

@JsonSerializable(explicitToJson: true)
class Address extends Equatable {
	@JsonKey(name: "street_line_1")final String street;
	@JsonKey(name: "City")final String city;
	@JsonKey(name: "State")final String state;
	
	const Address({
		required this.street,
		required this.city,
		required this.state,
	});
	
	Map<String, dynamic> toJson() => _$AddressToJson(this);
	
	factory Address.fromJson(Map<String, dynamic> json) => _$AddressFromJson(json);
	
	@override
	List<Object?> get props => [
		street,
		city,
		state,
	];
}

@JsonSerializable(explicitToJson: true)
class Profile extends Equatable {
	@JsonKey(name: "ID")final int id;
	@JsonKey(name: "full_name")final String name;
	@JsonKey(name: "Email")final String email;
	@JsonKey(name: "street_line_1")final String street;
	@JsonKey(name: "City")final String city;
	@JsonKey(name: "State")final String state;
	
	const Profile({
		required this.id,
		required this.name,
		required this.email,
		required this.street,
		required this.city,
		required this.state,
	});
	
	Map<String, dynamic> toJson() => _$ProfileToJson(this);
	
	factory Profile.fromJson(Map<String, dynamic> json) => _$ProfileFromJson(json);
	
	@override
	List<Object?> get props => [
		id,
		name,
		email,
		street,
		city,
		state,
	];
}

@JsonSerializable(explicitToJson: true)
class User extends Equatable {
	@JsonKey(name: "ID")final int id;
	@JsonKey(name: "full_name")final String name;
	@JsonKey(name: "Email")final String email;
	@JsonKey(name: "street_line_1")final String street;
	@JsonKey(name: "City")final String city;
	@JsonKey(name: "State")final String state;
	@JsonKey(name: "Password")final String password;
	@JsonKey(name: "CreatedAt")final DateTime createdAt;
	@JsonKey(name: "UpdatedAt")final DateTime updatedAt;
	@JsonKey(name: "DeletedAt")final DateTime? deletedAt;
	@JsonKey(defaultValue: <String, String>{}, name: "Options")final Map<String, String> options;
	@JsonKey(defaultValue: <List<String>>[], name: "Tags")final List<String> tags;
	
	const User({
		required this.id,
		required this.name,
		required this.email,
		required this.street,
		required this.city,
		required this.state,
		required this.password,
		required this.createdAt,
		required this.updatedAt,
		this.deletedAt,
		required this.options,
		required this.tags,
	});
	
	Map<String, dynamic> toJson() => _$UserToJson(this);
	
	factory User.fromJson(Map<String, dynamic> json) => _$UserFromJson(json);
	
	@override
	List<Object?> get props => [
		id,
		name,
		email,
		street,
		city,
		state,
		password,
		createdAt,
		updatedAt,
		deletedAt,
		options,
		tags,
	];
}

