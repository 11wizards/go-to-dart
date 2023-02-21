import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:json_annotation/json_annotation.dart';

part 'firestore.go.g.dart';

class _TimestampConverter implements JsonConverter<DateTime, Timestamp> {
  const _TimestampConverter();

  @override
  DateTime fromJson(Timestamp json) => json.toDate();

  @override
  Timestamp toJson(DateTime object) => Timestamp.fromDate(object);
}

@JsonSerializable()
@_TimestampConverter()
class User {
	final int id;
	final String name;
	final String email;
	final String password;
	final DateTime createdAt;
	final DateTime updatedAt;
	final DateTime? deletedAt;
	@JsonKey(defaultValue: <String, String>{})final Map<String, String> options;
	@JsonKey(defaultValue: <List<String>>[])final List<String> tags;
	
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

