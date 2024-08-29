// ignore_for_file: always_use_package_imports
import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:equatable/equatable.dart';
import 'package:json_annotation/json_annotation.dart';

part 'multipackage.go.g.dart';

class _TimestampConverter implements JsonConverter<DateTime, Timestamp> {
  const _TimestampConverter();

  @override
  DateTime fromJson(Timestamp json) => json.toDate();

  @override
  Timestamp toJson(DateTime object) => Timestamp.fromDate(object);
}

@JsonSerializable(explicitToJson: true)
@_TimestampConverter()
class Outer extends Equatable {
	final String? id;
	final String name;
	
	const Outer({
		this.id,
		required this.name,
	});
	
	Map<String, dynamic> toJson() => _$OuterToJson(this);
	
	factory Outer.fromJson(Map<String, dynamic> json) => _$OuterFromJson(json);
	
	@override
	List<Object?> get props => [
		id,
		name,
	];
}

