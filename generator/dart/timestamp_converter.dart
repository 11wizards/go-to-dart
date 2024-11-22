class _TimestampConverter implements JsonConverter<DateTime, Timestamp> {
  const _TimestampConverter();

  @override
  DateTime fromJson(Timestamp json) => DateTime.fromMicrosecondsSinceEpoch(json.microsecondsSinceEpoch, isUtc: true);

  @override
  Timestamp toJson(DateTime object) => Timestamp.fromDate(object.toUtc());
}