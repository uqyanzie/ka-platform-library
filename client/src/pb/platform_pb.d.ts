import * as jspb from 'google-protobuf'



export class NewPlatform extends jspb.Message {
  getUnitname(): string;
  setUnitname(value: string): NewPlatform;

  getUnitclass(): string;
  setUnitclass(value: string): NewPlatform;

  getOperationField(): string;
  setOperationField(value: string): NewPlatform;

  getGeneralcategory(): string;
  setGeneralcategory(value: string): NewPlatform;

  getGeneraltype(): string;
  setGeneraltype(value: string): NewPlatform;

  getLethalitylevel(): string;
  setLethalitylevel(value: string): NewPlatform;

  getCruisingspeed(): number;
  setCruisingspeed(value: number): NewPlatform;

  getMaxspeed(): number;
  setMaxspeed(value: number): NewPlatform;

  getCombatrange(): number;
  setCombatrange(value: number): NewPlatform;

  getFuelload(): number;
  setFuelload(value: number): NewPlatform;

  getCountryorigin(): string;
  setCountryorigin(value: string): NewPlatform;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NewPlatform.AsObject;
  static toObject(includeInstance: boolean, msg: NewPlatform): NewPlatform.AsObject;
  static serializeBinaryToWriter(message: NewPlatform, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NewPlatform;
  static deserializeBinaryFromReader(message: NewPlatform, reader: jspb.BinaryReader): NewPlatform;
}

export namespace NewPlatform {
  export type AsObject = {
    unitname: string,
    unitclass: string,
    operationField: string,
    generalcategory: string,
    generaltype: string,
    lethalitylevel: string,
    cruisingspeed: number,
    maxspeed: number,
    combatrange: number,
    fuelload: number,
    countryorigin: string,
  }
}

export class Platform extends jspb.Message {
  getUnitname(): string;
  setUnitname(value: string): Platform;

  getUnitclass(): string;
  setUnitclass(value: string): Platform;

  getOperationField(): string;
  setOperationField(value: string): Platform;

  getGeneralcategory(): string;
  setGeneralcategory(value: string): Platform;

  getGeneraltype(): string;
  setGeneraltype(value: string): Platform;

  getLethalitylevel(): string;
  setLethalitylevel(value: string): Platform;

  getCruisingspeed(): number;
  setCruisingspeed(value: number): Platform;

  getMaxspeed(): number;
  setMaxspeed(value: number): Platform;

  getCombatrange(): number;
  setCombatrange(value: number): Platform;

  getFuelload(): number;
  setFuelload(value: number): Platform;

  getCountryorigin(): string;
  setCountryorigin(value: string): Platform;

  getId(): string;
  setId(value: string): Platform;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Platform.AsObject;
  static toObject(includeInstance: boolean, msg: Platform): Platform.AsObject;
  static serializeBinaryToWriter(message: Platform, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Platform;
  static deserializeBinaryFromReader(message: Platform, reader: jspb.BinaryReader): Platform;
}

export namespace Platform {
  export type AsObject = {
    unitname: string,
    unitclass: string,
    operationField: string,
    generalcategory: string,
    generaltype: string,
    lethalitylevel: string,
    cruisingspeed: number,
    maxspeed: number,
    combatrange: number,
    fuelload: number,
    countryorigin: string,
    id: string,
  }
}

export class PlatformListRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PlatformListRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PlatformListRequest): PlatformListRequest.AsObject;
  static serializeBinaryToWriter(message: PlatformListRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlatformListRequest;
  static deserializeBinaryFromReader(message: PlatformListRequest, reader: jspb.BinaryReader): PlatformListRequest;
}

export namespace PlatformListRequest {
  export type AsObject = {
  }
}

export class PlatformListResponse extends jspb.Message {
  getPlatformsList(): Array<Platform>;
  setPlatformsList(value: Array<Platform>): PlatformListResponse;
  clearPlatformsList(): PlatformListResponse;
  addPlatforms(value?: Platform, index?: number): Platform;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PlatformListResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PlatformListResponse): PlatformListResponse.AsObject;
  static serializeBinaryToWriter(message: PlatformListResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlatformListResponse;
  static deserializeBinaryFromReader(message: PlatformListResponse, reader: jspb.BinaryReader): PlatformListResponse;
}

export namespace PlatformListResponse {
  export type AsObject = {
    platformsList: Array<Platform.AsObject>,
  }
}

