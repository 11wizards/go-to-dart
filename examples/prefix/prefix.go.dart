// ignore_for_file: always_use_package_imports
import 'package:copy_with_extension/copy_with_extension.dart';
import 'package:equatable/equatable.dart';
import 'package:json_annotation/json_annotation.dart';

part 'prefix.go.g.dart';

@CopyWith()
@JsonSerializable(explicitToJson: true)
class MyInstance extends Equatable {
	final MyMap<String, int> m;
	
	const MyInstance({
		required this.m,
	});
	
	Map<String, dynamic> toJson() => _$MyInstanceToJson(this);
	
	factory MyInstance.fromJson(Map<String, dynamic> json) => _$MyInstanceFromJson(json);
	
	@override
	List<Object?> get props => [
		m,
	];
}

@CopyWith()
@JsonSerializable(explicitToJson: true, genericArgumentFactories: true)
class MyKeyValuePair<TKey, TValue> extends Equatable {
	final TKey key;
	final TValue value;
	
	const MyKeyValuePair({
		required this.key,
		required this.value,
	});
	
	Map<String, dynamic> toJson(Object Function(TKey) toJsonTKey, Object Function(TValue) toJsonTValue) => _$MyKeyValuePairToJson(this, toJsonTKey, toJsonTValue);
	
	factory MyKeyValuePair.fromJson(Map<String, dynamic> json, TKey Function(Object? json) fromJsonTKey, TValue Function(Object? json) fromJsonTValue) => _$MyKeyValuePairFromJson(json, fromJsonTKey, fromJsonTValue);
	
	@override
	List<Object?> get props => [
		key,
		value,
	];
}

@CopyWith()
@JsonSerializable(explicitToJson: true, genericArgumentFactories: true)
class MyMap<TKey, TValue> extends Equatable {
	@JsonKey(defaultValue: <List<MyKeyValuePair<TKey, TValue>>>[])final List<MyKeyValuePair<TKey, TValue>> items;
	
	const MyMap({
		required this.items,
	});
	
	Map<String, dynamic> toJson(Object Function(TKey) toJsonTKey, Object Function(TValue) toJsonTValue) => _$MyMapToJson(this, toJsonTKey, toJsonTValue);
	
	factory MyMap.fromJson(Map<String, dynamic> json, TKey Function(Object? json) fromJsonTKey, TValue Function(Object? json) fromJsonTValue) => _$MyMapFromJson(json, fromJsonTKey, fromJsonTValue);
	
	@override
	List<Object?> get props => [
		items,
	];
}

@CopyWith()
@JsonSerializable(explicitToJson: true)
class MyUserRepository extends Equatable {
	@JsonKey(defaultValue: <List<User>>[])final List<User> users;
	
	const MyUserRepository({
		required this.users,
	});
	
	Map<String, dynamic> toJson() => _$MyUserRepositoryToJson(this);
	
	factory MyUserRepository.fromJson(Map<String, dynamic> json) => _$MyUserRepositoryFromJson(json);
	
	@override
	List<Object?> get props => [
		users,
	];
}

