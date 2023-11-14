import Vue from 'vue';

type FormatFunc = (string) => string;
type DetermineFunc = (string) => boolean;
declare module 'vue' {
  export interface ComponentCustomProperties {
    $formatUrl: FormatFunc;
    $logo: string;
    $parseTime: FormatFunc;
    $formatTime: FormatFunc;
    $densityList: Record<string, string>;
    $genderList: Record<string, string>;
    $hasPermission: DetermineFunc;
  }
}

declare module 'js-yaml';
