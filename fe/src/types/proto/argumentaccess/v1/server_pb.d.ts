// package: argumentaccess.v1
// file: argumentaccess/v1/server.proto

import * as jspb from "google-protobuf";
import * as google_api_annotations_pb from "../../google/api/annotations_pb";
import * as argumentaccess_v1_argument_pb from "../../argumentaccess/v1/argument_pb";

export class HealthCheckRequest extends jspb.Message {
  getService(): string;
  setService(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HealthCheckRequest.AsObject;
  static toObject(includeInstance: boolean, msg: HealthCheckRequest): HealthCheckRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HealthCheckRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HealthCheckRequest;
  static deserializeBinaryFromReader(message: HealthCheckRequest, reader: jspb.BinaryReader): HealthCheckRequest;
}

export namespace HealthCheckRequest {
  export type AsObject = {
    service: string,
  }
}

export class HealthCheckResponse extends jspb.Message {
  getStatus(): HealthCheckResponse.ServingStatusMap[keyof HealthCheckResponse.ServingStatusMap];
  setStatus(value: HealthCheckResponse.ServingStatusMap[keyof HealthCheckResponse.ServingStatusMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HealthCheckResponse.AsObject;
  static toObject(includeInstance: boolean, msg: HealthCheckResponse): HealthCheckResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HealthCheckResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HealthCheckResponse;
  static deserializeBinaryFromReader(message: HealthCheckResponse, reader: jspb.BinaryReader): HealthCheckResponse;
}

export namespace HealthCheckResponse {
  export type AsObject = {
    status: HealthCheckResponse.ServingStatusMap[keyof HealthCheckResponse.ServingStatusMap],
  }

  export interface ServingStatusMap {
    UNKNOWN: 0;
    SERVING: 1;
    NOT_SERVING: 2;
    SERVICE_UNKNOWN: 3;
  }

  export const ServingStatus: ServingStatusMap;
}

export class CreateSubjectRequest extends jspb.Message {
  hasSubject(): boolean;
  clearSubject(): void;
  getSubject(): argumentaccess_v1_argument_pb.Subject | undefined;
  setSubject(value?: argumentaccess_v1_argument_pb.Subject): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateSubjectRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateSubjectRequest): CreateSubjectRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateSubjectRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateSubjectRequest;
  static deserializeBinaryFromReader(message: CreateSubjectRequest, reader: jspb.BinaryReader): CreateSubjectRequest;
}

export namespace CreateSubjectRequest {
  export type AsObject = {
    subject?: argumentaccess_v1_argument_pb.Subject.AsObject,
  }
}

export class CreateSubjectResponse extends jspb.Message {
  hasSubject(): boolean;
  clearSubject(): void;
  getSubject(): argumentaccess_v1_argument_pb.Subject | undefined;
  setSubject(value?: argumentaccess_v1_argument_pb.Subject): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateSubjectResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateSubjectResponse): CreateSubjectResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateSubjectResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateSubjectResponse;
  static deserializeBinaryFromReader(message: CreateSubjectResponse, reader: jspb.BinaryReader): CreateSubjectResponse;
}

export namespace CreateSubjectResponse {
  export type AsObject = {
    subject?: argumentaccess_v1_argument_pb.Subject.AsObject,
  }
}

export class ReadSubjectRequest extends jspb.Message {
  hasSubjectId(): boolean;
  clearSubjectId(): void;
  getSubjectId(): string;
  setSubjectId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReadSubjectRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ReadSubjectRequest): ReadSubjectRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ReadSubjectRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReadSubjectRequest;
  static deserializeBinaryFromReader(message: ReadSubjectRequest, reader: jspb.BinaryReader): ReadSubjectRequest;
}

export namespace ReadSubjectRequest {
  export type AsObject = {
    subjectId: string,
  }
}

export class ReadSubjectResponse extends jspb.Message {
  clearSubjectsList(): void;
  getSubjectsList(): Array<argumentaccess_v1_argument_pb.Subject>;
  setSubjectsList(value: Array<argumentaccess_v1_argument_pb.Subject>): void;
  addSubjects(value?: argumentaccess_v1_argument_pb.Subject, index?: number): argumentaccess_v1_argument_pb.Subject;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReadSubjectResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ReadSubjectResponse): ReadSubjectResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ReadSubjectResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReadSubjectResponse;
  static deserializeBinaryFromReader(message: ReadSubjectResponse, reader: jspb.BinaryReader): ReadSubjectResponse;
}

export namespace ReadSubjectResponse {
  export type AsObject = {
    subjectsList: Array<argumentaccess_v1_argument_pb.Subject.AsObject>,
  }
}

export class UpdateSubjectRequest extends jspb.Message {
  hasSubject(): boolean;
  clearSubject(): void;
  getSubject(): argumentaccess_v1_argument_pb.Subject | undefined;
  setSubject(value?: argumentaccess_v1_argument_pb.Subject): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateSubjectRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateSubjectRequest): UpdateSubjectRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateSubjectRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateSubjectRequest;
  static deserializeBinaryFromReader(message: UpdateSubjectRequest, reader: jspb.BinaryReader): UpdateSubjectRequest;
}

export namespace UpdateSubjectRequest {
  export type AsObject = {
    subject?: argumentaccess_v1_argument_pb.Subject.AsObject,
  }
}

export class UpdateSubjectResponse extends jspb.Message {
  hasSubject(): boolean;
  clearSubject(): void;
  getSubject(): argumentaccess_v1_argument_pb.Subject | undefined;
  setSubject(value?: argumentaccess_v1_argument_pb.Subject): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateSubjectResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateSubjectResponse): UpdateSubjectResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateSubjectResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateSubjectResponse;
  static deserializeBinaryFromReader(message: UpdateSubjectResponse, reader: jspb.BinaryReader): UpdateSubjectResponse;
}

export namespace UpdateSubjectResponse {
  export type AsObject = {
    subject?: argumentaccess_v1_argument_pb.Subject.AsObject,
  }
}

export class DeleteSubjectRequest extends jspb.Message {
  getSubjectId(): string;
  setSubjectId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteSubjectRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteSubjectRequest): DeleteSubjectRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteSubjectRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteSubjectRequest;
  static deserializeBinaryFromReader(message: DeleteSubjectRequest, reader: jspb.BinaryReader): DeleteSubjectRequest;
}

export namespace DeleteSubjectRequest {
  export type AsObject = {
    subjectId: string,
  }
}

export class DeleteSubjectResponse extends jspb.Message {
  getDeletedCount(): number;
  setDeletedCount(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteSubjectResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteSubjectResponse): DeleteSubjectResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteSubjectResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteSubjectResponse;
  static deserializeBinaryFromReader(message: DeleteSubjectResponse, reader: jspb.BinaryReader): DeleteSubjectResponse;
}

export namespace DeleteSubjectResponse {
  export type AsObject = {
    deletedCount: number,
  }
}

export class CreatePredicateRequest extends jspb.Message {
  hasPredicate(): boolean;
  clearPredicate(): void;
  getPredicate(): argumentaccess_v1_argument_pb.Predicate | undefined;
  setPredicate(value?: argumentaccess_v1_argument_pb.Predicate): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreatePredicateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreatePredicateRequest): CreatePredicateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreatePredicateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreatePredicateRequest;
  static deserializeBinaryFromReader(message: CreatePredicateRequest, reader: jspb.BinaryReader): CreatePredicateRequest;
}

export namespace CreatePredicateRequest {
  export type AsObject = {
    predicate?: argumentaccess_v1_argument_pb.Predicate.AsObject,
  }
}

export class CreatePredicateResponse extends jspb.Message {
  hasPredicate(): boolean;
  clearPredicate(): void;
  getPredicate(): argumentaccess_v1_argument_pb.Predicate | undefined;
  setPredicate(value?: argumentaccess_v1_argument_pb.Predicate): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreatePredicateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreatePredicateResponse): CreatePredicateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreatePredicateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreatePredicateResponse;
  static deserializeBinaryFromReader(message: CreatePredicateResponse, reader: jspb.BinaryReader): CreatePredicateResponse;
}

export namespace CreatePredicateResponse {
  export type AsObject = {
    predicate?: argumentaccess_v1_argument_pb.Predicate.AsObject,
  }
}

export class ReadPredicateRequest extends jspb.Message {
  hasPredicateId(): boolean;
  clearPredicateId(): void;
  getPredicateId(): string;
  setPredicateId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReadPredicateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ReadPredicateRequest): ReadPredicateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ReadPredicateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReadPredicateRequest;
  static deserializeBinaryFromReader(message: ReadPredicateRequest, reader: jspb.BinaryReader): ReadPredicateRequest;
}

export namespace ReadPredicateRequest {
  export type AsObject = {
    predicateId: string,
  }
}

export class ReadPredicateResponse extends jspb.Message {
  clearPredicatesList(): void;
  getPredicatesList(): Array<argumentaccess_v1_argument_pb.Predicate>;
  setPredicatesList(value: Array<argumentaccess_v1_argument_pb.Predicate>): void;
  addPredicates(value?: argumentaccess_v1_argument_pb.Predicate, index?: number): argumentaccess_v1_argument_pb.Predicate;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReadPredicateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ReadPredicateResponse): ReadPredicateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ReadPredicateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReadPredicateResponse;
  static deserializeBinaryFromReader(message: ReadPredicateResponse, reader: jspb.BinaryReader): ReadPredicateResponse;
}

export namespace ReadPredicateResponse {
  export type AsObject = {
    predicatesList: Array<argumentaccess_v1_argument_pb.Predicate.AsObject>,
  }
}

export class UpdatePredicateRequest extends jspb.Message {
  hasPredicate(): boolean;
  clearPredicate(): void;
  getPredicate(): argumentaccess_v1_argument_pb.Predicate | undefined;
  setPredicate(value?: argumentaccess_v1_argument_pb.Predicate): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdatePredicateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdatePredicateRequest): UpdatePredicateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdatePredicateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdatePredicateRequest;
  static deserializeBinaryFromReader(message: UpdatePredicateRequest, reader: jspb.BinaryReader): UpdatePredicateRequest;
}

export namespace UpdatePredicateRequest {
  export type AsObject = {
    predicate?: argumentaccess_v1_argument_pb.Predicate.AsObject,
  }
}

export class UpdatePredicateResponse extends jspb.Message {
  hasPredicate(): boolean;
  clearPredicate(): void;
  getPredicate(): argumentaccess_v1_argument_pb.Predicate | undefined;
  setPredicate(value?: argumentaccess_v1_argument_pb.Predicate): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdatePredicateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdatePredicateResponse): UpdatePredicateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdatePredicateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdatePredicateResponse;
  static deserializeBinaryFromReader(message: UpdatePredicateResponse, reader: jspb.BinaryReader): UpdatePredicateResponse;
}

export namespace UpdatePredicateResponse {
  export type AsObject = {
    predicate?: argumentaccess_v1_argument_pb.Predicate.AsObject,
  }
}

export class DeletePredicateRequest extends jspb.Message {
  getPredicateId(): string;
  setPredicateId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeletePredicateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeletePredicateRequest): DeletePredicateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeletePredicateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeletePredicateRequest;
  static deserializeBinaryFromReader(message: DeletePredicateRequest, reader: jspb.BinaryReader): DeletePredicateRequest;
}

export namespace DeletePredicateRequest {
  export type AsObject = {
    predicateId: string,
  }
}

export class DeletePredicateResponse extends jspb.Message {
  getDeletedCount(): number;
  setDeletedCount(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeletePredicateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeletePredicateResponse): DeletePredicateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeletePredicateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeletePredicateResponse;
  static deserializeBinaryFromReader(message: DeletePredicateResponse, reader: jspb.BinaryReader): DeletePredicateResponse;
}

export namespace DeletePredicateResponse {
  export type AsObject = {
    deletedCount: number,
  }
}

export class CreatePremiseRequest extends jspb.Message {
  hasPremise(): boolean;
  clearPremise(): void;
  getPremise(): argumentaccess_v1_argument_pb.Premise | undefined;
  setPremise(value?: argumentaccess_v1_argument_pb.Premise): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreatePremiseRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreatePremiseRequest): CreatePremiseRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreatePremiseRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreatePremiseRequest;
  static deserializeBinaryFromReader(message: CreatePremiseRequest, reader: jspb.BinaryReader): CreatePremiseRequest;
}

export namespace CreatePremiseRequest {
  export type AsObject = {
    premise?: argumentaccess_v1_argument_pb.Premise.AsObject,
  }
}

export class CreatePremiseResponse extends jspb.Message {
  hasPremise(): boolean;
  clearPremise(): void;
  getPremise(): argumentaccess_v1_argument_pb.Premise | undefined;
  setPremise(value?: argumentaccess_v1_argument_pb.Premise): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreatePremiseResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreatePremiseResponse): CreatePremiseResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreatePremiseResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreatePremiseResponse;
  static deserializeBinaryFromReader(message: CreatePremiseResponse, reader: jspb.BinaryReader): CreatePremiseResponse;
}

export namespace CreatePremiseResponse {
  export type AsObject = {
    premise?: argumentaccess_v1_argument_pb.Premise.AsObject,
  }
}

export class ReadPremiseRequest extends jspb.Message {
  getPremiseId(): string;
  setPremiseId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReadPremiseRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ReadPremiseRequest): ReadPremiseRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ReadPremiseRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReadPremiseRequest;
  static deserializeBinaryFromReader(message: ReadPremiseRequest, reader: jspb.BinaryReader): ReadPremiseRequest;
}

export namespace ReadPremiseRequest {
  export type AsObject = {
    premiseId: string,
  }
}

export class ReadPremiseResponse extends jspb.Message {
  clearPremisesList(): void;
  getPremisesList(): Array<argumentaccess_v1_argument_pb.Premise>;
  setPremisesList(value: Array<argumentaccess_v1_argument_pb.Premise>): void;
  addPremises(value?: argumentaccess_v1_argument_pb.Premise, index?: number): argumentaccess_v1_argument_pb.Premise;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReadPremiseResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ReadPremiseResponse): ReadPremiseResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ReadPremiseResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReadPremiseResponse;
  static deserializeBinaryFromReader(message: ReadPremiseResponse, reader: jspb.BinaryReader): ReadPremiseResponse;
}

export namespace ReadPremiseResponse {
  export type AsObject = {
    premisesList: Array<argumentaccess_v1_argument_pb.Premise.AsObject>,
  }
}

export class UpdatePremiseRequest extends jspb.Message {
  hasPremise(): boolean;
  clearPremise(): void;
  getPremise(): argumentaccess_v1_argument_pb.Premise | undefined;
  setPremise(value?: argumentaccess_v1_argument_pb.Premise): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdatePremiseRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdatePremiseRequest): UpdatePremiseRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdatePremiseRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdatePremiseRequest;
  static deserializeBinaryFromReader(message: UpdatePremiseRequest, reader: jspb.BinaryReader): UpdatePremiseRequest;
}

export namespace UpdatePremiseRequest {
  export type AsObject = {
    premise?: argumentaccess_v1_argument_pb.Premise.AsObject,
  }
}

export class UpdatePremiseResponse extends jspb.Message {
  hasPremise(): boolean;
  clearPremise(): void;
  getPremise(): argumentaccess_v1_argument_pb.Premise | undefined;
  setPremise(value?: argumentaccess_v1_argument_pb.Premise): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdatePremiseResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdatePremiseResponse): UpdatePremiseResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdatePremiseResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdatePremiseResponse;
  static deserializeBinaryFromReader(message: UpdatePremiseResponse, reader: jspb.BinaryReader): UpdatePremiseResponse;
}

export namespace UpdatePremiseResponse {
  export type AsObject = {
    premise?: argumentaccess_v1_argument_pb.Premise.AsObject,
  }
}

export class DeletePremiseRequest extends jspb.Message {
  getPremiseId(): string;
  setPremiseId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeletePremiseRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeletePremiseRequest): DeletePremiseRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeletePremiseRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeletePremiseRequest;
  static deserializeBinaryFromReader(message: DeletePremiseRequest, reader: jspb.BinaryReader): DeletePremiseRequest;
}

export namespace DeletePremiseRequest {
  export type AsObject = {
    premiseId: string,
  }
}

export class DeletePremiseResponse extends jspb.Message {
  getDeletedCount(): number;
  setDeletedCount(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeletePremiseResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeletePremiseResponse): DeletePremiseResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeletePremiseResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeletePremiseResponse;
  static deserializeBinaryFromReader(message: DeletePremiseResponse, reader: jspb.BinaryReader): DeletePremiseResponse;
}

export namespace DeletePremiseResponse {
  export type AsObject = {
    deletedCount: number,
  }
}

export class CreatePropositionRequest extends jspb.Message {
  hasProposition(): boolean;
  clearProposition(): void;
  getProposition(): argumentaccess_v1_argument_pb.Proposition | undefined;
  setProposition(value?: argumentaccess_v1_argument_pb.Proposition): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreatePropositionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreatePropositionRequest): CreatePropositionRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreatePropositionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreatePropositionRequest;
  static deserializeBinaryFromReader(message: CreatePropositionRequest, reader: jspb.BinaryReader): CreatePropositionRequest;
}

export namespace CreatePropositionRequest {
  export type AsObject = {
    proposition?: argumentaccess_v1_argument_pb.Proposition.AsObject,
  }
}

export class CreatePropositionResponse extends jspb.Message {
  hasProposition(): boolean;
  clearProposition(): void;
  getProposition(): argumentaccess_v1_argument_pb.Proposition | undefined;
  setProposition(value?: argumentaccess_v1_argument_pb.Proposition): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreatePropositionResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreatePropositionResponse): CreatePropositionResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreatePropositionResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreatePropositionResponse;
  static deserializeBinaryFromReader(message: CreatePropositionResponse, reader: jspb.BinaryReader): CreatePropositionResponse;
}

export namespace CreatePropositionResponse {
  export type AsObject = {
    proposition?: argumentaccess_v1_argument_pb.Proposition.AsObject,
  }
}

export class ReadPropositionRequest extends jspb.Message {
  getPropositionId(): string;
  setPropositionId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReadPropositionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ReadPropositionRequest): ReadPropositionRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ReadPropositionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReadPropositionRequest;
  static deserializeBinaryFromReader(message: ReadPropositionRequest, reader: jspb.BinaryReader): ReadPropositionRequest;
}

export namespace ReadPropositionRequest {
  export type AsObject = {
    propositionId: string,
  }
}

export class ReadPropositionResponse extends jspb.Message {
  clearPropositionsList(): void;
  getPropositionsList(): Array<argumentaccess_v1_argument_pb.Proposition>;
  setPropositionsList(value: Array<argumentaccess_v1_argument_pb.Proposition>): void;
  addPropositions(value?: argumentaccess_v1_argument_pb.Proposition, index?: number): argumentaccess_v1_argument_pb.Proposition;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReadPropositionResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ReadPropositionResponse): ReadPropositionResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ReadPropositionResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReadPropositionResponse;
  static deserializeBinaryFromReader(message: ReadPropositionResponse, reader: jspb.BinaryReader): ReadPropositionResponse;
}

export namespace ReadPropositionResponse {
  export type AsObject = {
    propositionsList: Array<argumentaccess_v1_argument_pb.Proposition.AsObject>,
  }
}

export class UpdatePropositionRequest extends jspb.Message {
  hasProposition(): boolean;
  clearProposition(): void;
  getProposition(): argumentaccess_v1_argument_pb.Proposition | undefined;
  setProposition(value?: argumentaccess_v1_argument_pb.Proposition): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdatePropositionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdatePropositionRequest): UpdatePropositionRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdatePropositionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdatePropositionRequest;
  static deserializeBinaryFromReader(message: UpdatePropositionRequest, reader: jspb.BinaryReader): UpdatePropositionRequest;
}

export namespace UpdatePropositionRequest {
  export type AsObject = {
    proposition?: argumentaccess_v1_argument_pb.Proposition.AsObject,
  }
}

export class UpdatePropositionResponse extends jspb.Message {
  hasProposition(): boolean;
  clearProposition(): void;
  getProposition(): argumentaccess_v1_argument_pb.Proposition | undefined;
  setProposition(value?: argumentaccess_v1_argument_pb.Proposition): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdatePropositionResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdatePropositionResponse): UpdatePropositionResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdatePropositionResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdatePropositionResponse;
  static deserializeBinaryFromReader(message: UpdatePropositionResponse, reader: jspb.BinaryReader): UpdatePropositionResponse;
}

export namespace UpdatePropositionResponse {
  export type AsObject = {
    proposition?: argumentaccess_v1_argument_pb.Proposition.AsObject,
  }
}

export class DeletePropositionRequest extends jspb.Message {
  getPropositionId(): string;
  setPropositionId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeletePropositionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeletePropositionRequest): DeletePropositionRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeletePropositionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeletePropositionRequest;
  static deserializeBinaryFromReader(message: DeletePropositionRequest, reader: jspb.BinaryReader): DeletePropositionRequest;
}

export namespace DeletePropositionRequest {
  export type AsObject = {
    propositionId: string,
  }
}

export class DeletePropositionResponse extends jspb.Message {
  getDeletedCount(): number;
  setDeletedCount(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeletePropositionResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeletePropositionResponse): DeletePropositionResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeletePropositionResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeletePropositionResponse;
  static deserializeBinaryFromReader(message: DeletePropositionResponse, reader: jspb.BinaryReader): DeletePropositionResponse;
}

export namespace DeletePropositionResponse {
  export type AsObject = {
    deletedCount: number,
  }
}

export class CreateConditionalStatementRequest extends jspb.Message {
  hasConditionalStatement(): boolean;
  clearConditionalStatement(): void;
  getConditionalStatement(): argumentaccess_v1_argument_pb.ConditionalStatement | undefined;
  setConditionalStatement(value?: argumentaccess_v1_argument_pb.ConditionalStatement): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateConditionalStatementRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateConditionalStatementRequest): CreateConditionalStatementRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateConditionalStatementRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateConditionalStatementRequest;
  static deserializeBinaryFromReader(message: CreateConditionalStatementRequest, reader: jspb.BinaryReader): CreateConditionalStatementRequest;
}

export namespace CreateConditionalStatementRequest {
  export type AsObject = {
    conditionalStatement?: argumentaccess_v1_argument_pb.ConditionalStatement.AsObject,
  }
}

export class CreateConditionalStatementResponse extends jspb.Message {
  hasConditionalStatement(): boolean;
  clearConditionalStatement(): void;
  getConditionalStatement(): argumentaccess_v1_argument_pb.ConditionalStatement | undefined;
  setConditionalStatement(value?: argumentaccess_v1_argument_pb.ConditionalStatement): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateConditionalStatementResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateConditionalStatementResponse): CreateConditionalStatementResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateConditionalStatementResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateConditionalStatementResponse;
  static deserializeBinaryFromReader(message: CreateConditionalStatementResponse, reader: jspb.BinaryReader): CreateConditionalStatementResponse;
}

export namespace CreateConditionalStatementResponse {
  export type AsObject = {
    conditionalStatement?: argumentaccess_v1_argument_pb.ConditionalStatement.AsObject,
  }
}

export class ReadConditionalStatementRequest extends jspb.Message {
  getConditionalStatementId(): string;
  setConditionalStatementId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReadConditionalStatementRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ReadConditionalStatementRequest): ReadConditionalStatementRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ReadConditionalStatementRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReadConditionalStatementRequest;
  static deserializeBinaryFromReader(message: ReadConditionalStatementRequest, reader: jspb.BinaryReader): ReadConditionalStatementRequest;
}

export namespace ReadConditionalStatementRequest {
  export type AsObject = {
    conditionalStatementId: string,
  }
}

export class ReadConditionalStatementResponse extends jspb.Message {
  clearConditionalStatementsList(): void;
  getConditionalStatementsList(): Array<argumentaccess_v1_argument_pb.ConditionalStatement>;
  setConditionalStatementsList(value: Array<argumentaccess_v1_argument_pb.ConditionalStatement>): void;
  addConditionalStatements(value?: argumentaccess_v1_argument_pb.ConditionalStatement, index?: number): argumentaccess_v1_argument_pb.ConditionalStatement;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReadConditionalStatementResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ReadConditionalStatementResponse): ReadConditionalStatementResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ReadConditionalStatementResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReadConditionalStatementResponse;
  static deserializeBinaryFromReader(message: ReadConditionalStatementResponse, reader: jspb.BinaryReader): ReadConditionalStatementResponse;
}

export namespace ReadConditionalStatementResponse {
  export type AsObject = {
    conditionalStatementsList: Array<argumentaccess_v1_argument_pb.ConditionalStatement.AsObject>,
  }
}

export class UpdateConditionalStatementRequest extends jspb.Message {
  hasConditionalStatement(): boolean;
  clearConditionalStatement(): void;
  getConditionalStatement(): argumentaccess_v1_argument_pb.ConditionalStatement | undefined;
  setConditionalStatement(value?: argumentaccess_v1_argument_pb.ConditionalStatement): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateConditionalStatementRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateConditionalStatementRequest): UpdateConditionalStatementRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateConditionalStatementRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateConditionalStatementRequest;
  static deserializeBinaryFromReader(message: UpdateConditionalStatementRequest, reader: jspb.BinaryReader): UpdateConditionalStatementRequest;
}

export namespace UpdateConditionalStatementRequest {
  export type AsObject = {
    conditionalStatement?: argumentaccess_v1_argument_pb.ConditionalStatement.AsObject,
  }
}

export class UpdateConditionalStatementResponse extends jspb.Message {
  hasConditionalStatement(): boolean;
  clearConditionalStatement(): void;
  getConditionalStatement(): argumentaccess_v1_argument_pb.ConditionalStatement | undefined;
  setConditionalStatement(value?: argumentaccess_v1_argument_pb.ConditionalStatement): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateConditionalStatementResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateConditionalStatementResponse): UpdateConditionalStatementResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateConditionalStatementResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateConditionalStatementResponse;
  static deserializeBinaryFromReader(message: UpdateConditionalStatementResponse, reader: jspb.BinaryReader): UpdateConditionalStatementResponse;
}

export namespace UpdateConditionalStatementResponse {
  export type AsObject = {
    conditionalStatement?: argumentaccess_v1_argument_pb.ConditionalStatement.AsObject,
  }
}

export class DeleteConditionalStatementRequest extends jspb.Message {
  getConditionalStatementId(): string;
  setConditionalStatementId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteConditionalStatementRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteConditionalStatementRequest): DeleteConditionalStatementRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteConditionalStatementRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteConditionalStatementRequest;
  static deserializeBinaryFromReader(message: DeleteConditionalStatementRequest, reader: jspb.BinaryReader): DeleteConditionalStatementRequest;
}

export namespace DeleteConditionalStatementRequest {
  export type AsObject = {
    conditionalStatementId: string,
  }
}

export class DeleteConditionalStatementResponse extends jspb.Message {
  getDeletedCount(): number;
  setDeletedCount(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteConditionalStatementResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteConditionalStatementResponse): DeleteConditionalStatementResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteConditionalStatementResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteConditionalStatementResponse;
  static deserializeBinaryFromReader(message: DeleteConditionalStatementResponse, reader: jspb.BinaryReader): DeleteConditionalStatementResponse;
}

export namespace DeleteConditionalStatementResponse {
  export type AsObject = {
    deletedCount: number,
  }
}

export class CreateArgumentRequest extends jspb.Message {
  hasArgument(): boolean;
  clearArgument(): void;
  getArgument(): argumentaccess_v1_argument_pb.Argument | undefined;
  setArgument(value?: argumentaccess_v1_argument_pb.Argument): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateArgumentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateArgumentRequest): CreateArgumentRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateArgumentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateArgumentRequest;
  static deserializeBinaryFromReader(message: CreateArgumentRequest, reader: jspb.BinaryReader): CreateArgumentRequest;
}

export namespace CreateArgumentRequest {
  export type AsObject = {
    argument?: argumentaccess_v1_argument_pb.Argument.AsObject,
  }
}

export class CreateArgumentResponse extends jspb.Message {
  hasArgument(): boolean;
  clearArgument(): void;
  getArgument(): argumentaccess_v1_argument_pb.Argument | undefined;
  setArgument(value?: argumentaccess_v1_argument_pb.Argument): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateArgumentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateArgumentResponse): CreateArgumentResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateArgumentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateArgumentResponse;
  static deserializeBinaryFromReader(message: CreateArgumentResponse, reader: jspb.BinaryReader): CreateArgumentResponse;
}

export namespace CreateArgumentResponse {
  export type AsObject = {
    argument?: argumentaccess_v1_argument_pb.Argument.AsObject,
  }
}

export class ReadArgumentRequest extends jspb.Message {
  getArgumentId(): string;
  setArgumentId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReadArgumentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ReadArgumentRequest): ReadArgumentRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ReadArgumentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReadArgumentRequest;
  static deserializeBinaryFromReader(message: ReadArgumentRequest, reader: jspb.BinaryReader): ReadArgumentRequest;
}

export namespace ReadArgumentRequest {
  export type AsObject = {
    argumentId: string,
  }
}

export class ReadArgumentResponse extends jspb.Message {
  clearArgumentsList(): void;
  getArgumentsList(): Array<argumentaccess_v1_argument_pb.Argument>;
  setArgumentsList(value: Array<argumentaccess_v1_argument_pb.Argument>): void;
  addArguments(value?: argumentaccess_v1_argument_pb.Argument, index?: number): argumentaccess_v1_argument_pb.Argument;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReadArgumentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ReadArgumentResponse): ReadArgumentResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ReadArgumentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReadArgumentResponse;
  static deserializeBinaryFromReader(message: ReadArgumentResponse, reader: jspb.BinaryReader): ReadArgumentResponse;
}

export namespace ReadArgumentResponse {
  export type AsObject = {
    argumentsList: Array<argumentaccess_v1_argument_pb.Argument.AsObject>,
  }
}

export class UpdateArgumentRequest extends jspb.Message {
  hasArgument(): boolean;
  clearArgument(): void;
  getArgument(): argumentaccess_v1_argument_pb.Argument | undefined;
  setArgument(value?: argumentaccess_v1_argument_pb.Argument): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateArgumentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateArgumentRequest): UpdateArgumentRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateArgumentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateArgumentRequest;
  static deserializeBinaryFromReader(message: UpdateArgumentRequest, reader: jspb.BinaryReader): UpdateArgumentRequest;
}

export namespace UpdateArgumentRequest {
  export type AsObject = {
    argument?: argumentaccess_v1_argument_pb.Argument.AsObject,
  }
}

export class UpdateArgumentResponse extends jspb.Message {
  hasArgument(): boolean;
  clearArgument(): void;
  getArgument(): argumentaccess_v1_argument_pb.Argument | undefined;
  setArgument(value?: argumentaccess_v1_argument_pb.Argument): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateArgumentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateArgumentResponse): UpdateArgumentResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateArgumentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateArgumentResponse;
  static deserializeBinaryFromReader(message: UpdateArgumentResponse, reader: jspb.BinaryReader): UpdateArgumentResponse;
}

export namespace UpdateArgumentResponse {
  export type AsObject = {
    argument?: argumentaccess_v1_argument_pb.Argument.AsObject,
  }
}

export class DeleteArgumentRequest extends jspb.Message {
  getArgumentId(): string;
  setArgumentId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteArgumentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteArgumentRequest): DeleteArgumentRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteArgumentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteArgumentRequest;
  static deserializeBinaryFromReader(message: DeleteArgumentRequest, reader: jspb.BinaryReader): DeleteArgumentRequest;
}

export namespace DeleteArgumentRequest {
  export type AsObject = {
    argumentId: string,
  }
}

export class DeleteArgumentResponse extends jspb.Message {
  getDeletedCount(): number;
  setDeletedCount(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteArgumentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteArgumentResponse): DeleteArgumentResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteArgumentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteArgumentResponse;
  static deserializeBinaryFromReader(message: DeleteArgumentResponse, reader: jspb.BinaryReader): DeleteArgumentResponse;
}

export namespace DeleteArgumentResponse {
  export type AsObject = {
    deletedCount: number,
  }
}

export class TestRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TestRequest.AsObject;
  static toObject(includeInstance: boolean, msg: TestRequest): TestRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TestRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TestRequest;
  static deserializeBinaryFromReader(message: TestRequest, reader: jspb.BinaryReader): TestRequest;
}

export namespace TestRequest {
  export type AsObject = {
  }
}

export class TestResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TestResponse.AsObject;
  static toObject(includeInstance: boolean, msg: TestResponse): TestResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TestResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TestResponse;
  static deserializeBinaryFromReader(message: TestResponse, reader: jspb.BinaryReader): TestResponse;
}

export namespace TestResponse {
  export type AsObject = {
  }
}

