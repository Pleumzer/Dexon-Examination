import { CmlsInterface } from "./ICml";

export interface InfosInterface {
  ID?: number ;
  Line_Number?: string;
  Location?: string;
  From?: string;
  To?: string;
  DrawingNumber?: string;
  Service?: string;
  Material?: string;
  InService_Date?: Date;
  PipeSize?: number;
  OriginalThickness?: number;
  Stress?: number;
  Joint_Efficiency?: number;
  CA?: number;
  Design_Life?: number;
  Design_Pressure?: number;
  Operating_Pressure?: number;
  Design_Temperature?: number;
  Operating_Temperature?: number;

  Cmls?: CmlsInterface[];
}
