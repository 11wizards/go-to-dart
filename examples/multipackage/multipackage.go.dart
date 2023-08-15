import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:json_annotation/json_annotation.dart';

part 'multipackage.go.g.dart';

class _TimestampConverter implements JsonConverter<DateTime, Timestamp> {
  const _TimestampConverter();

  @override
  DateTime fromJson(Timestamp json) => json.toDate();

  @override
  Timestamp toJson(DateTime object) => Timestamp.fromDate(object);
}

@JsonSerializable()
@_TimestampConverter()
class Outer {
	final String? id;
	final String name;
	
	Outer({
		this.id,
		required this.name,
	});
	
	Map<String, dynamic> toJson() => _$OuterToJson(this);
	
	factory Outer.fromJson(Map<String, dynamic> json) => _$OuterFromJson(json);
}

