// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file rill/admin/v1/ai.proto (package rill.admin.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { CompletionMessage, Tool } from "../../ai/v1/ai_pb.js";

/**
 * @generated from message rill.admin.v1.CompleteRequest
 */
export class CompleteRequest extends Message<CompleteRequest> {
  /**
   * @generated from field: repeated rill.ai.v1.CompletionMessage messages = 1;
   */
  messages: CompletionMessage[] = [];

  /**
   * Optional list of tools that the AI can use during completion
   *
   * @generated from field: repeated rill.ai.v1.Tool tools = 2;
   */
  tools: Tool[] = [];

  constructor(data?: PartialMessage<CompleteRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "rill.admin.v1.CompleteRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "messages", kind: "message", T: CompletionMessage, repeated: true },
    { no: 2, name: "tools", kind: "message", T: Tool, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CompleteRequest {
    return new CompleteRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CompleteRequest {
    return new CompleteRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CompleteRequest {
    return new CompleteRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CompleteRequest | PlainMessage<CompleteRequest> | undefined, b: CompleteRequest | PlainMessage<CompleteRequest> | undefined): boolean {
    return proto3.util.equals(CompleteRequest, a, b);
  }
}

/**
 * @generated from message rill.admin.v1.CompleteResponse
 */
export class CompleteResponse extends Message<CompleteResponse> {
  /**
   * @generated from field: rill.ai.v1.CompletionMessage message = 1;
   */
  message?: CompletionMessage;

  constructor(data?: PartialMessage<CompleteResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "rill.admin.v1.CompleteResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "message", kind: "message", T: CompletionMessage },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CompleteResponse {
    return new CompleteResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CompleteResponse {
    return new CompleteResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CompleteResponse {
    return new CompleteResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CompleteResponse | PlainMessage<CompleteResponse> | undefined, b: CompleteResponse | PlainMessage<CompleteResponse> | undefined): boolean {
    return proto3.util.equals(CompleteResponse, a, b);
  }
}

