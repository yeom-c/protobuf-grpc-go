// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: protos/character_collection/character_collection.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace GrpcCharacterCollection {

  /// <summary>Holder for reflection information generated from protos/character_collection/character_collection.proto</summary>
  public static partial class CharacterCollectionReflection {

    #region Descriptor
    /// <summary>File descriptor for protos/character_collection/character_collection.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static CharacterCollectionReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "CjZwcm90b3MvY2hhcmFjdGVyX2NvbGxlY3Rpb24vY2hhcmFjdGVyX2NvbGxl",
            "Y3Rpb24ucHJvdG8SFGNoYXJhY3Rlcl9jb2xsZWN0aW9uGhhwcm90b3MvbW9k",
            "ZWwvbW9kZWwucHJvdG8iWwoaR2V0Q2hhcmFjdGVyQ29sbGVjdGlvbnNSZXMS",
            "PQoVY2hhcmFjdGVyX2NvbGxlY3Rpb25zGAEgAygLMh4ubW9kZWwuQ2hhcmFj",
            "dGVyQ29sbGVjdGlvbkRhdGEyeQoaQ2hhcmFjdGVyQ29sbGVjdGlvblNlcnZp",
            "Y2USWwoXR2V0Q2hhcmFjdGVyQ29sbGVjdGlvbnMSDC5tb2RlbC5FbXB0eRow",
            "LmNoYXJhY3Rlcl9jb2xsZWN0aW9uLkdldENoYXJhY3RlckNvbGxlY3Rpb25z",
            "UmVzIgBCZVpJZ2l0aHViLmNvbS95ZW9tLWMvcHJvdG9idWYtZ3JwYy1nby9n",
            "ZW4vZ29sYW5nL3Byb3Rvcy9jaGFyYWN0ZXJfY29sbGVjdGlvbqoCF0dycGND",
            "aGFyYWN0ZXJDb2xsZWN0aW9uYgZwcm90bzM="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::GrpcModel.ModelReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::GrpcCharacterCollection.GetCharacterCollectionsRes), global::GrpcCharacterCollection.GetCharacterCollectionsRes.Parser, new[]{ "CharacterCollections" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class GetCharacterCollectionsRes : pb::IMessage<GetCharacterCollectionsRes>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<GetCharacterCollectionsRes> _parser = new pb::MessageParser<GetCharacterCollectionsRes>(() => new GetCharacterCollectionsRes());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<GetCharacterCollectionsRes> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::GrpcCharacterCollection.CharacterCollectionReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public GetCharacterCollectionsRes() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public GetCharacterCollectionsRes(GetCharacterCollectionsRes other) : this() {
      characterCollections_ = other.characterCollections_.Clone();
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public GetCharacterCollectionsRes Clone() {
      return new GetCharacterCollectionsRes(this);
    }

    /// <summary>Field number for the "character_collections" field.</summary>
    public const int CharacterCollectionsFieldNumber = 1;
    private static readonly pb::FieldCodec<global::GrpcModel.CharacterCollectionData> _repeated_characterCollections_codec
        = pb::FieldCodec.ForMessage(10, global::GrpcModel.CharacterCollectionData.Parser);
    private readonly pbc::RepeatedField<global::GrpcModel.CharacterCollectionData> characterCollections_ = new pbc::RepeatedField<global::GrpcModel.CharacterCollectionData>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public pbc::RepeatedField<global::GrpcModel.CharacterCollectionData> CharacterCollections {
      get { return characterCollections_; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as GetCharacterCollectionsRes);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(GetCharacterCollectionsRes other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if(!characterCollections_.Equals(other.characterCollections_)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      hash ^= characterCollections_.GetHashCode();
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
      characterCollections_.WriteTo(output, _repeated_characterCollections_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      characterCollections_.WriteTo(ref output, _repeated_characterCollections_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int CalculateSize() {
      int size = 0;
      size += characterCollections_.CalculateSize(_repeated_characterCollections_codec);
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(GetCharacterCollectionsRes other) {
      if (other == null) {
        return;
      }
      characterCollections_.Add(other.characterCollections_);
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
            characterCollections_.AddEntriesFrom(input, _repeated_characterCollections_codec);
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
            characterCollections_.AddEntriesFrom(ref input, _repeated_characterCollections_codec);
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
