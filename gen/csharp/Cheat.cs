// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: protos/cheat/cheat.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace GrpcCheat {

  /// <summary>Holder for reflection information generated from protos/cheat/cheat.proto</summary>
  public static partial class CheatReflection {

    #region Descriptor
    /// <summary>File descriptor for protos/cheat/cheat.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static CheatReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "Chhwcm90b3MvY2hlYXQvY2hlYXQucHJvdG8SBWNoZWF0Ghhwcm90b3MvbW9k",
            "ZWwvbW9kZWwucHJvdG8iMQoOQ3JlYXRlQXNzZXRSZXESDwoHZW51bV9pZBgB",
            "IAEoCRIOCgZhbW91bnQYAiABKAMycwoMQ2hlYXRTZXJ2aWNlEiwKC0NyZWF0",
            "ZU1haWxzEgwubW9kZWwuRW1wdHkaDS5tb2RlbC5SZXN1bHQiABI1CgtDcmVh",
            "dGVBc3NldBIVLmNoZWF0LkNyZWF0ZUFzc2V0UmVxGg0ubW9kZWwuUmVzdWx0",
            "IgBCSFo6Z2l0aHViLmNvbS95ZW9tLWMvcHJvdG9idWYtZ3JwYy1nby9nZW4v",
            "Z29sYW5nL3Byb3Rvcy9jaGVhdKoCCUdycGNDaGVhdGIGcHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::GrpcModel.ModelReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::GrpcCheat.CreateAssetReq), global::GrpcCheat.CreateAssetReq.Parser, new[]{ "EnumId", "Amount" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class CreateAssetReq : pb::IMessage<CreateAssetReq>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<CreateAssetReq> _parser = new pb::MessageParser<CreateAssetReq>(() => new CreateAssetReq());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<CreateAssetReq> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::GrpcCheat.CheatReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public CreateAssetReq() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public CreateAssetReq(CreateAssetReq other) : this() {
      enumId_ = other.enumId_;
      amount_ = other.amount_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public CreateAssetReq Clone() {
      return new CreateAssetReq(this);
    }

    /// <summary>Field number for the "enum_id" field.</summary>
    public const int EnumIdFieldNumber = 1;
    private string enumId_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public string EnumId {
      get { return enumId_; }
      set {
        enumId_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "amount" field.</summary>
    public const int AmountFieldNumber = 2;
    private long amount_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public long Amount {
      get { return amount_; }
      set {
        amount_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as CreateAssetReq);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(CreateAssetReq other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (EnumId != other.EnumId) return false;
      if (Amount != other.Amount) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (EnumId.Length != 0) hash ^= EnumId.GetHashCode();
      if (Amount != 0L) hash ^= Amount.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      if (EnumId.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(EnumId);
      }
      if (Amount != 0L) {
        output.WriteRawTag(16);
        output.WriteInt64(Amount);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (EnumId.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(EnumId);
      }
      if (Amount != 0L) {
        output.WriteRawTag(16);
        output.WriteInt64(Amount);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int CalculateSize() {
      int size = 0;
      if (EnumId.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(EnumId);
      }
      if (Amount != 0L) {
        size += 1 + pb::CodedOutputStream.ComputeInt64Size(Amount);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(CreateAssetReq other) {
      if (other == null) {
        return;
      }
      if (other.EnumId.Length != 0) {
        EnumId = other.EnumId;
      }
      if (other.Amount != 0L) {
        Amount = other.Amount;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            EnumId = input.ReadString();
            break;
          }
          case 16: {
            Amount = input.ReadInt64();
            break;
          }
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
          case 10: {
            EnumId = input.ReadString();
            break;
          }
          case 16: {
            Amount = input.ReadInt64();
            break;
          }
        }
      }
    }
    #endif

  }

  #endregion

}

#endregion Designer generated code