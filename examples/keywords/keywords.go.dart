// ignore_for_file: always_use_package_imports
import 'package:copy_with_extension/copy_with_extension.dart';
import 'package:equatable/equatable.dart';
import 'package:json_annotation/json_annotation.dart';

part 'keywords.go.g.dart';

@CopyWith()
@JsonSerializable(explicitToJson: true)
class Class$ extends Equatable {
	final int id;
	final String name;
	
	const Class$({
		required this.id,
		required this.name,
	});
	
	Map<String, dynamic> toJson() => _$Class$ToJson(this);
	
	factory Class$.fromJson(Map<String, dynamic> json) => _$Class$FromJson(json);
	
	@override
	List<Object?> get props => [
		id,
		name,
	];
}

@CopyWith()
@JsonSerializable(explicitToJson: true)
class Interface$ extends Equatable {
	final int id;
	@JsonKey(name: "type")final String type$;
	
	const Interface$({
		required this.id,
		required this.type$,
	});
	
	Map<String, dynamic> toJson() => _$Interface$ToJson(this);
	
	factory Interface$.fromJson(Map<String, dynamic> json) => _$Interface$FromJson(json);
	
	@override
	List<Object?> get props => [
		id,
		type$,
	];
}

@CopyWith()
@JsonSerializable(explicitToJson: true)
class TestKeywords extends Equatable {
	final int id;
	@JsonKey(name: "class")final String class$;
	@JsonKey(name: "static")final String static$;
	@JsonKey(name: "final")final String final$;
	@JsonKey(name: "const")final String const$;
	@JsonKey(name: "var")final String var$;
	@JsonKey(name: "if")final String if$;
	@JsonKey(name: "else")final String else$;
	@JsonKey(name: "for")final String for$;
	@JsonKey(name: "while")final String while$;
	@JsonKey(name: "switch")final String switch$;
	@JsonKey(name: "case")final String case$;
	@JsonKey(name: "default")final String default$;
	@JsonKey(name: "break")final String break$;
	@JsonKey(name: "continue")final String continue$;
	@JsonKey(name: "return")final String return$;
	@JsonKey(name: "try")final String try$;
	@JsonKey(name: "catch")final String catch$;
	@JsonKey(name: "finally")final String finally$;
	@JsonKey(name: "throw")final String throw$;
	@JsonKey(name: "new")final String new$;
	@JsonKey(name: "this")final String this$;
	@JsonKey(name: "super")final String super$;
	@JsonKey(name: "null")final String null$;
	@JsonKey(name: "true")final String true$;
	@JsonKey(name: "false")final String false$;
	@JsonKey(name: "async")final String async$;
	@JsonKey(name: "await")final String await$;
	@JsonKey(name: "yield")final String yield$;
	@JsonKey(name: "abstract")final String abstract$;
	@JsonKey(name: "extends")final String extends$;
	@JsonKey(name: "with")final String with$;
	@JsonKey(name: "mixin")final String mixin$;
	@JsonKey(name: "enum")final String enum$;
	final DateTime createdAt;
	
	const TestKeywords({
		required this.id,
		required this.class$,
		required this.static$,
		required this.final$,
		required this.const$,
		required this.var$,
		required this.if$,
		required this.else$,
		required this.for$,
		required this.while$,
		required this.switch$,
		required this.case$,
		required this.default$,
		required this.break$,
		required this.continue$,
		required this.return$,
		required this.try$,
		required this.catch$,
		required this.finally$,
		required this.throw$,
		required this.new$,
		required this.this$,
		required this.super$,
		required this.null$,
		required this.true$,
		required this.false$,
		required this.async$,
		required this.await$,
		required this.yield$,
		required this.abstract$,
		required this.extends$,
		required this.with$,
		required this.mixin$,
		required this.enum$,
		required this.createdAt,
	});
	
	Map<String, dynamic> toJson() => _$TestKeywordsToJson(this);
	
	factory TestKeywords.fromJson(Map<String, dynamic> json) => _$TestKeywordsFromJson(json);
	
	@override
	List<Object?> get props => [
		id,
		class$,
		static$,
		final$,
		const$,
		var$,
		if$,
		else$,
		for$,
		while$,
		switch$,
		case$,
		default$,
		break$,
		continue$,
		return$,
		try$,
		catch$,
		finally$,
		throw$,
		new$,
		this$,
		super$,
		null$,
		true$,
		false$,
		async$,
		await$,
		yield$,
		abstract$,
		extends$,
		with$,
		mixin$,
		enum$,
		createdAt,
	];
}

