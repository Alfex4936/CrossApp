import { main } from "../models/models"

export function Parse(url: string, length: number): Promise<Array<main.Notice>>;

export function GetWeather(): Promise<main.Weather>;

export function GetPeople(keyword: string): Promise<Array<main.People>>;
