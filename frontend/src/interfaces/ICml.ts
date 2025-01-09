import { TestPointsInterface } from "./ITestPoint";

export interface CmlsInterface{
    ID?: number;

  Cml_number?: number;

  Cml_description?: string;

  Actual_Outside_Diameter?: number;

  Design_thickness?: number;

  Structural_thickness?: number;

  Required_thickness?: number;

  InfoID?: number;

  TestPoints?: TestPointsInterface[];
}