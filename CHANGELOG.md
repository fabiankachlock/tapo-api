## v2.0.0

BREAKING CHANGES

`github.com/fabiankachlock/tapo-api/api/requests`:

- Renamed `RequestGetTriggerLog` to `RequestGetTriggerLogs`
- Renamed `AlarmVolumeMedium` to `AlarmVolumeNormal`
- `PlayAlarmParams` must be used via the builder methods how.
- `EnergyDataParams` must be used via the builder methods how.



`github.com/fabiankachlock/tapo-api/api`:

- The api client has been completely reworked. Please refer to the documentation for the new usage.