// ignore_for_file: always_use_package_imports
import 'package:json_annotation/json_annotation.dart';

part 'generics.go.g.dart';

@JsonSerializable(explicitToJson: true)
class Instance {
	final Map<String, int> m;
	
	Instance({
		required this.m,
	});
	
	Map<String, dynamic> toJson() => _$InstanceToJson(this);
	
	factory Instance.fromJson(Map<String, dynamic> json) => _$InstanceFromJson(json);
}

@JsonSerializable(explicitToJson: true, genericArgumentFactories: true)
class KeyValuePair<TKey, TValue> {
	final TKey key;
	final TValue value;
	
	KeyValuePair({
		required this.key,
		required this.value,
	});
	
	Map<String, dynamic> toJson(Object Function(TKey) toJsonTKey, Object Function(TValue) toJsonTValue) => _$KeyValuePairToJson(this, toJsonTKey, toJsonTValue);
	
	factory KeyValuePair.fromJson(Map<String, dynamic> json, TKey Function(Object? json) fromJsonTKey, TValue Function(Object? json) fromJsonTValue) => _$KeyValuePairFromJson(json, fromJsonTKey, fromJsonTValue);
}

@JsonSerializable(explicitToJson: true, genericArgumentFactories: true)
class Map<TKey, TValue> {
	@JsonKey(defaultValue: <List<KeyValuePair<TKey, TValue>>>[])final List<KeyValuePair<TKey, TValue>> items;
	
	Map({
		required this.items,
	});
	
	Map<String, dynamic> toJson(Object Function(TKey) toJsonTKey, Object Function(TValue) toJsonTValue) => _$MapToJson(this, toJsonTKey, toJsonTValue);
	
	factory Map.fromJson(Map<String, dynamic> json, TKey Function(Object? json) fromJsonTKey, TValue Function(Object? json) fromJsonTValue) => _$MapFromJson(json, fromJsonTKey, fromJsonTValue);
}

