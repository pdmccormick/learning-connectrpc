// @generated by protoc-gen-es v1.3.0 with parameter "target=ts"
// @generated from file chat/v1/chat.proto (package chat.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type {
	BinaryReadOptions,
	FieldList,
	JsonReadOptions,
	JsonValue,
	PartialMessage,
	PlainMessage
} from '@bufbuild/protobuf';
import { Message, proto3 } from '@bufbuild/protobuf';

/**
 * Say
 *
 * @generated from message chat.v1.SayRequest
 */
export class SayRequest extends Message<SayRequest> {
	/**
	 * @generated from field: string message = 1;
	 */
	message = '';

	constructor(data?: PartialMessage<SayRequest>) {
		super();
		proto3.util.initPartial(data, this);
	}

	static readonly runtime: typeof proto3 = proto3;
	static readonly typeName = 'chat.v1.SayRequest';
	static readonly fields: FieldList = proto3.util.newFieldList(() => [
		{ no: 1, name: 'message', kind: 'scalar', T: 9 /* ScalarType.STRING */ }
	]);

	static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SayRequest {
		return new SayRequest().fromBinary(bytes, options);
	}

	static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SayRequest {
		return new SayRequest().fromJson(jsonValue, options);
	}

	static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SayRequest {
		return new SayRequest().fromJsonString(jsonString, options);
	}

	static equals(
		a: SayRequest | PlainMessage<SayRequest> | undefined,
		b: SayRequest | PlainMessage<SayRequest> | undefined
	): boolean {
		return proto3.util.equals(SayRequest, a, b);
	}
}

/**
 * @generated from message chat.v1.SayResponse
 */
export class SayResponse extends Message<SayResponse> {
	/**
	 * @generated from field: bool ok = 1;
	 */
	ok = false;

	/**
	 * @generated from field: string error = 2;
	 */
	error = '';

	constructor(data?: PartialMessage<SayResponse>) {
		super();
		proto3.util.initPartial(data, this);
	}

	static readonly runtime: typeof proto3 = proto3;
	static readonly typeName = 'chat.v1.SayResponse';
	static readonly fields: FieldList = proto3.util.newFieldList(() => [
		{ no: 1, name: 'ok', kind: 'scalar', T: 8 /* ScalarType.BOOL */ },
		{ no: 2, name: 'error', kind: 'scalar', T: 9 /* ScalarType.STRING */ }
	]);

	static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SayResponse {
		return new SayResponse().fromBinary(bytes, options);
	}

	static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SayResponse {
		return new SayResponse().fromJson(jsonValue, options);
	}

	static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SayResponse {
		return new SayResponse().fromJsonString(jsonString, options);
	}

	static equals(
		a: SayResponse | PlainMessage<SayResponse> | undefined,
		b: SayResponse | PlainMessage<SayResponse> | undefined
	): boolean {
		return proto3.util.equals(SayResponse, a, b);
	}
}

/**
 * Listen
 *
 * @generated from message chat.v1.ListenRequest
 */
export class ListenRequest extends Message<ListenRequest> {
	constructor(data?: PartialMessage<ListenRequest>) {
		super();
		proto3.util.initPartial(data, this);
	}

	static readonly runtime: typeof proto3 = proto3;
	static readonly typeName = 'chat.v1.ListenRequest';
	static readonly fields: FieldList = proto3.util.newFieldList(() => []);

	static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListenRequest {
		return new ListenRequest().fromBinary(bytes, options);
	}

	static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListenRequest {
		return new ListenRequest().fromJson(jsonValue, options);
	}

	static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListenRequest {
		return new ListenRequest().fromJsonString(jsonString, options);
	}

	static equals(
		a: ListenRequest | PlainMessage<ListenRequest> | undefined,
		b: ListenRequest | PlainMessage<ListenRequest> | undefined
	): boolean {
		return proto3.util.equals(ListenRequest, a, b);
	}
}

/**
 * @generated from message chat.v1.ListenResponse
 */
export class ListenResponse extends Message<ListenResponse> {
	/**
	 * @generated from field: string person = 1;
	 */
	person = '';

	/**
	 * @generated from field: string message = 2;
	 */
	message = '';

	constructor(data?: PartialMessage<ListenResponse>) {
		super();
		proto3.util.initPartial(data, this);
	}

	static readonly runtime: typeof proto3 = proto3;
	static readonly typeName = 'chat.v1.ListenResponse';
	static readonly fields: FieldList = proto3.util.newFieldList(() => [
		{ no: 1, name: 'person', kind: 'scalar', T: 9 /* ScalarType.STRING */ },
		{ no: 2, name: 'message', kind: 'scalar', T: 9 /* ScalarType.STRING */ }
	]);

	static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListenResponse {
		return new ListenResponse().fromBinary(bytes, options);
	}

	static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListenResponse {
		return new ListenResponse().fromJson(jsonValue, options);
	}

	static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListenResponse {
		return new ListenResponse().fromJsonString(jsonString, options);
	}

	static equals(
		a: ListenResponse | PlainMessage<ListenResponse> | undefined,
		b: ListenResponse | PlainMessage<ListenResponse> | undefined
	): boolean {
		return proto3.util.equals(ListenResponse, a, b);
	}
}
