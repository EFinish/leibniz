// package: argumentaccess.v1
// file: argumentaccess/v1/argument.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class Argument extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  hasCreatedAt(): boolean;
  clearCreatedAt(): void;
  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasUpdatedAt(): boolean;
  clearUpdatedAt(): void;
  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasDeletedAt(): boolean;
  clearDeletedAt(): void;
  getDeletedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDeletedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getTitle(): string;
  setTitle(value: string): void;

  clearPremisesList(): void;
  getPremisesList(): Array<Premise>;
  setPremisesList(value: Array<Premise>): void;
  addPremises(value?: Premise, index?: number): Premise;

  clearConditionalStatementsList(): void;
  getConditionalStatementsList(): Array<ConditionalStatement>;
  setConditionalStatementsList(value: Array<ConditionalStatement>): void;
  addConditionalStatements(value?: ConditionalStatement, index?: number): ConditionalStatement;

  hasConclusionPremise(): boolean;
  clearConclusionPremise(): void;
  getConclusionPremise(): Premise | undefined;
  setConclusionPremise(value?: Premise): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Argument.AsObject;
  static toObject(includeInstance: boolean, msg: Argument): Argument.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Argument, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Argument;
  static deserializeBinaryFromReader(message: Argument, reader: jspb.BinaryReader): Argument;
}

export namespace Argument {
  export type AsObject = {
    id: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    deletedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    title: string,
    premisesList: Array<Premise.AsObject>,
    conditionalStatementsList: Array<ConditionalStatement.AsObject>,
    conclusionPremise?: Premise.AsObject,
  }
}

export class ConditionalStatement extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  hasCreatedAt(): boolean;
  clearCreatedAt(): void;
  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasUpdatedAt(): boolean;
  clearUpdatedAt(): void;
  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasDeletedAt(): boolean;
  clearDeletedAt(): void;
  getDeletedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDeletedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasIfPremise(): boolean;
  clearIfPremise(): void;
  getIfPremise(): Premise | undefined;
  setIfPremise(value?: Premise): void;

  hasThenPremise(): boolean;
  clearThenPremise(): void;
  getThenPremise(): Premise | undefined;
  setThenPremise(value?: Premise): void;

  hasIfProposition(): boolean;
  clearIfProposition(): void;
  getIfProposition(): Proposition | undefined;
  setIfProposition(value?: Proposition): void;

  hasThenProposition(): boolean;
  clearThenProposition(): void;
  getThenProposition(): Proposition | undefined;
  setThenProposition(value?: Proposition): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ConditionalStatement.AsObject;
  static toObject(includeInstance: boolean, msg: ConditionalStatement): ConditionalStatement.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ConditionalStatement, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ConditionalStatement;
  static deserializeBinaryFromReader(message: ConditionalStatement, reader: jspb.BinaryReader): ConditionalStatement;
}

export namespace ConditionalStatement {
  export type AsObject = {
    id: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    deletedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    ifPremise?: Premise.AsObject,
    thenPremise?: Premise.AsObject,
    ifProposition?: Proposition.AsObject,
    thenProposition?: Proposition.AsObject,
  }
}

export class Proposition extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  hasCreatedAt(): boolean;
  clearCreatedAt(): void;
  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasUpdatedAt(): boolean;
  clearUpdatedAt(): void;
  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasDeletedAt(): boolean;
  clearDeletedAt(): void;
  getDeletedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDeletedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getPropositionType(): PropositionTypeMap[keyof PropositionTypeMap];
  setPropositionType(value: PropositionTypeMap[keyof PropositionTypeMap]): void;

  clearPropositionSubPremisesList(): void;
  getPropositionSubPremisesList(): Array<Premise>;
  setPropositionSubPremisesList(value: Array<Premise>): void;
  addPropositionSubPremises(value?: Premise, index?: number): Premise;

  clearPropositionSubPropositionsList(): void;
  getPropositionSubPropositionsList(): Array<Proposition>;
  setPropositionSubPropositionsList(value: Array<Proposition>): void;
  addPropositionSubPropositions(value?: Proposition, index?: number): Proposition;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Proposition.AsObject;
  static toObject(includeInstance: boolean, msg: Proposition): Proposition.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Proposition, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Proposition;
  static deserializeBinaryFromReader(message: Proposition, reader: jspb.BinaryReader): Proposition;
}

export namespace Proposition {
  export type AsObject = {
    id: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    deletedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    propositionType: PropositionTypeMap[keyof PropositionTypeMap],
    propositionSubPremisesList: Array<Premise.AsObject>,
    propositionSubPropositionsList: Array<Proposition.AsObject>,
  }
}

export class Premise extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  hasCreatedAt(): boolean;
  clearCreatedAt(): void;
  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasUpdatedAt(): boolean;
  clearUpdatedAt(): void;
  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasDeletedAt(): boolean;
  clearDeletedAt(): void;
  getDeletedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDeletedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasSubject(): boolean;
  clearSubject(): void;
  getSubject(): Subject | undefined;
  setSubject(value?: Subject): void;

  getSubjectQuantifier(): SubjectQuantifierMap[keyof SubjectQuantifierMap];
  setSubjectQuantifier(value: SubjectQuantifierMap[keyof SubjectQuantifierMap]): void;

  hasPredicate(): boolean;
  clearPredicate(): void;
  getPredicate(): Predicate | undefined;
  setPredicate(value?: Predicate): void;

  getPredicateQualifier(): PredicateQualifierMap[keyof PredicateQualifierMap];
  setPredicateQualifier(value: PredicateQualifierMap[keyof PredicateQualifierMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Premise.AsObject;
  static toObject(includeInstance: boolean, msg: Premise): Premise.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Premise, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Premise;
  static deserializeBinaryFromReader(message: Premise, reader: jspb.BinaryReader): Premise;
}

export namespace Premise {
  export type AsObject = {
    id: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    deletedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    subject?: Subject.AsObject,
    subjectQuantifier: SubjectQuantifierMap[keyof SubjectQuantifierMap],
    predicate?: Predicate.AsObject,
    predicateQualifier: PredicateQualifierMap[keyof PredicateQualifierMap],
  }
}

export class Subject extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  hasCreatedAt(): boolean;
  clearCreatedAt(): void;
  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasUpdatedAt(): boolean;
  clearUpdatedAt(): void;
  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasDeletedAt(): boolean;
  clearDeletedAt(): void;
  getDeletedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDeletedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getBody(): string;
  setBody(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Subject.AsObject;
  static toObject(includeInstance: boolean, msg: Subject): Subject.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Subject, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Subject;
  static deserializeBinaryFromReader(message: Subject, reader: jspb.BinaryReader): Subject;
}

export namespace Subject {
  export type AsObject = {
    id: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    deletedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    body: string,
  }
}

export class Predicate extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  hasCreatedAt(): boolean;
  clearCreatedAt(): void;
  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasUpdatedAt(): boolean;
  clearUpdatedAt(): void;
  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasDeletedAt(): boolean;
  clearDeletedAt(): void;
  getDeletedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDeletedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getBody(): string;
  setBody(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Predicate.AsObject;
  static toObject(includeInstance: boolean, msg: Predicate): Predicate.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Predicate, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Predicate;
  static deserializeBinaryFromReader(message: Predicate, reader: jspb.BinaryReader): Predicate;
}

export namespace Predicate {
  export type AsObject = {
    id: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    deletedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    body: string,
  }
}

export interface PropositionTypeMap {
  AND: 0;
  OR: 1;
  NAND: 2;
  NOR: 3;
  XOR: 4;
  XNOR: 5;
}

export const PropositionType: PropositionTypeMap;

export interface SubjectQuantifierMap {
  ALL: 0;
  SOME: 1;
  NONE: 2;
}

export const SubjectQuantifier: SubjectQuantifierMap;

export interface PredicateQualifierMap {
  NOTISAREHASHAVE: 0;
  ISAREHASHAVE: 1;
}

export const PredicateQualifier: PredicateQualifierMap;

