### cram-go

Quantization and compression of scalars and basic geometry types.

For a .NET-version, check out [cram-dotnet](https://github.com/Piot/cram-dotnet).

Uses [basal-go](https://github.com/Piot/basal-go) for basic geometry types (e.g. Vector and Quaternion) and [brook-go](https://github.com/Piot/brook-go) for bitstreams.

##### Usage

###### Structure
```go
q := types.NewQuaternion(-0.183, 0.683, -0.062, 0.704)
packedInfo := QuaternionPack(&q)
q2, err := QuaternionUnPack(packedInfo)
```

###### Stream

```go
cramStream := inbitstream.New(brookInBitStream)
q, err := cramStream.ReadRotation()
```

