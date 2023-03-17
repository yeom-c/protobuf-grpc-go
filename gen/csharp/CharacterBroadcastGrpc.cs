// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: protos/character_broadcast/character_broadcast.proto
// </auto-generated>
#pragma warning disable 0414, 1591, 8981
#region Designer generated code

using grpc = global::Grpc.Core;

namespace GrpcCharacterBroadcast {
  public static partial class CharacterBroadcastService
  {
    static readonly string __ServiceName = "character_broadcast.CharacterBroadcastService";

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static void __Helper_SerializeMessage(global::Google.Protobuf.IMessage message, grpc::SerializationContext context)
    {
      #if !GRPC_DISABLE_PROTOBUF_BUFFER_SERIALIZATION
      if (message is global::Google.Protobuf.IBufferMessage)
      {
        context.SetPayloadLength(message.CalculateSize());
        global::Google.Protobuf.MessageExtensions.WriteTo(message, context.GetBufferWriter());
        context.Complete();
        return;
      }
      #endif
      context.Complete(global::Google.Protobuf.MessageExtensions.ToByteArray(message));
    }

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static class __Helper_MessageCache<T>
    {
      public static readonly bool IsBufferMessage = global::System.Reflection.IntrospectionExtensions.GetTypeInfo(typeof(global::Google.Protobuf.IBufferMessage)).IsAssignableFrom(typeof(T));
    }

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static T __Helper_DeserializeMessage<T>(grpc::DeserializationContext context, global::Google.Protobuf.MessageParser<T> parser) where T : global::Google.Protobuf.IMessage<T>
    {
      #if !GRPC_DISABLE_PROTOBUF_BUFFER_SERIALIZATION
      if (__Helper_MessageCache<T>.IsBufferMessage)
      {
        return parser.ParseFrom(context.PayloadAsReadOnlySequence());
      }
      #endif
      return parser.ParseFrom(context.PayloadAsNewBuffer());
    }

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::GrpcModel.Empty> __Marshaller_model_Empty = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::GrpcModel.Empty.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::GrpcCharacterBroadcast.GetOnAirCharacterBroadcastsRes> __Marshaller_character_broadcast_GetOnAirCharacterBroadcastsRes = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::GrpcCharacterBroadcast.GetOnAirCharacterBroadcastsRes.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::GrpcCharacterBroadcast.GetCompletedCharacterBroadcastsReq> __Marshaller_character_broadcast_GetCompletedCharacterBroadcastsReq = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::GrpcCharacterBroadcast.GetCompletedCharacterBroadcastsReq.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::GrpcCharacterBroadcast.GetCompletedCharacterBroadcastsRes> __Marshaller_character_broadcast_GetCompletedCharacterBroadcastsRes = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::GrpcCharacterBroadcast.GetCompletedCharacterBroadcastsRes.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::GrpcCharacterBroadcast.CompleteCharacterBroadcastReq> __Marshaller_character_broadcast_CompleteCharacterBroadcastReq = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::GrpcCharacterBroadcast.CompleteCharacterBroadcastReq.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::GrpcCharacterBroadcast.CompleteCharacterBroadcastRes> __Marshaller_character_broadcast_CompleteCharacterBroadcastRes = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::GrpcCharacterBroadcast.CompleteCharacterBroadcastRes.Parser));

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::GrpcModel.Empty, global::GrpcCharacterBroadcast.GetOnAirCharacterBroadcastsRes> __Method_GetOnAirCharacterBroadcasts = new grpc::Method<global::GrpcModel.Empty, global::GrpcCharacterBroadcast.GetOnAirCharacterBroadcastsRes>(
        grpc::MethodType.Unary,
        __ServiceName,
        "GetOnAirCharacterBroadcasts",
        __Marshaller_model_Empty,
        __Marshaller_character_broadcast_GetOnAirCharacterBroadcastsRes);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::GrpcCharacterBroadcast.GetCompletedCharacterBroadcastsReq, global::GrpcCharacterBroadcast.GetCompletedCharacterBroadcastsRes> __Method_GetCompletedCharacterBroadcasts = new grpc::Method<global::GrpcCharacterBroadcast.GetCompletedCharacterBroadcastsReq, global::GrpcCharacterBroadcast.GetCompletedCharacterBroadcastsRes>(
        grpc::MethodType.Unary,
        __ServiceName,
        "GetCompletedCharacterBroadcasts",
        __Marshaller_character_broadcast_GetCompletedCharacterBroadcastsReq,
        __Marshaller_character_broadcast_GetCompletedCharacterBroadcastsRes);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::GrpcCharacterBroadcast.CompleteCharacterBroadcastReq, global::GrpcCharacterBroadcast.CompleteCharacterBroadcastRes> __Method_CompleteCharacterBroadcast = new grpc::Method<global::GrpcCharacterBroadcast.CompleteCharacterBroadcastReq, global::GrpcCharacterBroadcast.CompleteCharacterBroadcastRes>(
        grpc::MethodType.Unary,
        __ServiceName,
        "CompleteCharacterBroadcast",
        __Marshaller_character_broadcast_CompleteCharacterBroadcastReq,
        __Marshaller_character_broadcast_CompleteCharacterBroadcastRes);

    /// <summary>Service descriptor</summary>
    public static global::Google.Protobuf.Reflection.ServiceDescriptor Descriptor
    {
      get { return global::GrpcCharacterBroadcast.CharacterBroadcastReflection.Descriptor.Services[0]; }
    }

    /// <summary>Client for CharacterBroadcastService</summary>
    public partial class CharacterBroadcastServiceClient : grpc::ClientBase<CharacterBroadcastServiceClient>
    {
      /// <summary>Creates a new client for CharacterBroadcastService</summary>
      /// <param name="channel">The channel to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public CharacterBroadcastServiceClient(grpc::ChannelBase channel) : base(channel)
      {
      }
      /// <summary>Creates a new client for CharacterBroadcastService that uses a custom <c>CallInvoker</c>.</summary>
      /// <param name="callInvoker">The callInvoker to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public CharacterBroadcastServiceClient(grpc::CallInvoker callInvoker) : base(callInvoker)
      {
      }
      /// <summary>Protected parameterless constructor to allow creation of test doubles.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected CharacterBroadcastServiceClient() : base()
      {
      }
      /// <summary>Protected constructor to allow creation of configured clients.</summary>
      /// <param name="configuration">The client configuration.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected CharacterBroadcastServiceClient(ClientBaseConfiguration configuration) : base(configuration)
      {
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::GrpcCharacterBroadcast.GetOnAirCharacterBroadcastsRes GetOnAirCharacterBroadcasts(global::GrpcModel.Empty request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetOnAirCharacterBroadcasts(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::GrpcCharacterBroadcast.GetOnAirCharacterBroadcastsRes GetOnAirCharacterBroadcasts(global::GrpcModel.Empty request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_GetOnAirCharacterBroadcasts, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::GrpcCharacterBroadcast.GetOnAirCharacterBroadcastsRes> GetOnAirCharacterBroadcastsAsync(global::GrpcModel.Empty request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetOnAirCharacterBroadcastsAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::GrpcCharacterBroadcast.GetOnAirCharacterBroadcastsRes> GetOnAirCharacterBroadcastsAsync(global::GrpcModel.Empty request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_GetOnAirCharacterBroadcasts, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::GrpcCharacterBroadcast.GetCompletedCharacterBroadcastsRes GetCompletedCharacterBroadcasts(global::GrpcCharacterBroadcast.GetCompletedCharacterBroadcastsReq request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetCompletedCharacterBroadcasts(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::GrpcCharacterBroadcast.GetCompletedCharacterBroadcastsRes GetCompletedCharacterBroadcasts(global::GrpcCharacterBroadcast.GetCompletedCharacterBroadcastsReq request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_GetCompletedCharacterBroadcasts, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::GrpcCharacterBroadcast.GetCompletedCharacterBroadcastsRes> GetCompletedCharacterBroadcastsAsync(global::GrpcCharacterBroadcast.GetCompletedCharacterBroadcastsReq request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetCompletedCharacterBroadcastsAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::GrpcCharacterBroadcast.GetCompletedCharacterBroadcastsRes> GetCompletedCharacterBroadcastsAsync(global::GrpcCharacterBroadcast.GetCompletedCharacterBroadcastsReq request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_GetCompletedCharacterBroadcasts, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::GrpcCharacterBroadcast.CompleteCharacterBroadcastRes CompleteCharacterBroadcast(global::GrpcCharacterBroadcast.CompleteCharacterBroadcastReq request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CompleteCharacterBroadcast(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::GrpcCharacterBroadcast.CompleteCharacterBroadcastRes CompleteCharacterBroadcast(global::GrpcCharacterBroadcast.CompleteCharacterBroadcastReq request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_CompleteCharacterBroadcast, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::GrpcCharacterBroadcast.CompleteCharacterBroadcastRes> CompleteCharacterBroadcastAsync(global::GrpcCharacterBroadcast.CompleteCharacterBroadcastReq request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CompleteCharacterBroadcastAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::GrpcCharacterBroadcast.CompleteCharacterBroadcastRes> CompleteCharacterBroadcastAsync(global::GrpcCharacterBroadcast.CompleteCharacterBroadcastReq request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_CompleteCharacterBroadcast, null, options, request);
      }
      /// <summary>Creates a new instance of client from given <c>ClientBaseConfiguration</c>.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected override CharacterBroadcastServiceClient NewInstance(ClientBaseConfiguration configuration)
      {
        return new CharacterBroadcastServiceClient(configuration);
      }
    }

  }
}
#endregion