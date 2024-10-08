// ignore_for_file: always_use_package_imports
import 'package:copy_with_extension/copy_with_extension.dart';
import 'package:equatable/equatable.dart';
import 'package:json_annotation/json_annotation.dart';

part 'everything.go.g.dart';

@CopyWith()
@JsonSerializable(explicitToJson: true)
class Child extends Equatable {
	final int id;
	final String name;
	
	const Child({
		required this.id,
		required this.name,
	});
	
	Map<String, dynamic> toJson() => _$ChildToJson(this);
	
	factory Child.fromJson(Map<String, dynamic> json) => _$ChildFromJson(json);
	
	@override
	List<Object?> get props => [
		id,
		name,
	];
}

@JsonSerializable(explicitToJson: true)
class Empty extends Equatable {
	
	const Empty();
	
	Map<String, dynamic> toJson() => _$EmptyToJson(this);
	
	factory Empty.fromJson(Map<String, dynamic> json) => _$EmptyFromJson(json);
	
	@override
	List<Object?> get props => [];
}

@CopyWith()
@JsonSerializable(explicitToJson: true)
class Parent extends Equatable {
	final String id;
	final int number1;
	@JsonKey(defaultValue: <List<int>>[])final List<int> number2;
	final int? number3;
	final List<int>? number4;
	@JsonKey(defaultValue: <List<int?>>[])final List<int?> number5;
	final int? number6;
	@JsonKey(defaultValue: <List<int?>>[])final List<int?> number7;
	final String text1;
	@JsonKey(defaultValue: <List<String>>[])final List<String> text2;
	final String? text3;
	final List<String>? text4;
	@JsonKey(defaultValue: <List<String?>>[])final List<String?> text5;
	final String? text6;
	@JsonKey(defaultValue: <List<String?>>[])final List<String?> text7;
	final DateTime date1;
	@JsonKey(defaultValue: <List<DateTime>>[])final List<DateTime> date2;
	final DateTime? date3;
	final List<DateTime>? date4;
	@JsonKey(defaultValue: <List<DateTime?>>[])final List<DateTime?> date5;
	final String? date6;
	@JsonKey(defaultValue: <List<String?>>[])final List<String?> date7;
	final Child child1;
	@JsonKey(defaultValue: <List<Child>>[])final List<Child> child2;
	final Child? child3;
	final List<Child>? child4;
	@JsonKey(defaultValue: <List<Child?>>[])final List<Child?> child5;
	final Child? child6;
	@JsonKey(defaultValue: <List<Child?>>[])final List<Child?> child7;
	@JsonKey(defaultValue: <String, double>{})final Map<String, double> map1;
	@JsonKey(defaultValue: <int, Child>{}, name: "map2_weird_name")final Map<int, Child> map2;
	@JsonKey(name: "empty")final Empty empty1;
	
	const Parent({
		required this.id,
		required this.number1,
		required this.number2,
		this.number3,
		this.number4,
		required this.number5,
		this.number6,
		required this.number7,
		required this.text1,
		required this.text2,
		this.text3,
		this.text4,
		required this.text5,
		this.text6,
		required this.text7,
		required this.date1,
		required this.date2,
		this.date3,
		this.date4,
		required this.date5,
		this.date6,
		required this.date7,
		required this.child1,
		required this.child2,
		this.child3,
		this.child4,
		required this.child5,
		this.child6,
		required this.child7,
		required this.map1,
		required this.map2,
		required this.empty1,
	});
	
	Map<String, dynamic> toJson() => _$ParentToJson(this);
	
	factory Parent.fromJson(Map<String, dynamic> json) => _$ParentFromJson(json);
	
	@override
	List<Object?> get props => [
		id,
		number1,
		number2,
		number3,
		number4,
		number5,
		number6,
		number7,
		text1,
		text2,
		text3,
		text4,
		text5,
		text6,
		text7,
		date1,
		date2,
		date3,
		date4,
		date5,
		date6,
		date7,
		child1,
		child2,
		child3,
		child4,
		child5,
		child6,
		child7,
		map1,
		map2,
		empty1,
	];
}

