export namespace main {
    export class Notice {
        id: number;
        category: string;
        title: string;
        date: string;
        link: string;
        writer: string;

        static createFrom(source: any = {}) {
            return new Notice(source);
        }

        constructor(source: any = {}) {
            if ('string' === typeof source) source = JSON.parse(source);
            this.id = source["id"];
            this.category = source["category"];
            this.title = source["title"];
            this.date = source["date"];
            this.link = source["link"];
            this.writer = source["writer"];
        }
    }
    export class Weather {
        max_temp: string;
        min_temp: string;
        current_temp: string;
        current_stat: string;
        rain_day: string;
        rain_night: string;
        fine_dust: string;
        ultra_dust: string;
        uv: string;
        sunset: string;
        icon: string;
        tmrw_rain_day: string;
        tmrw_rain_night: string;

        static createFrom(source: any = {}) {
            return new Weather(source);
        }

        constructor(source: any = {}) {
            if ('string' === typeof source) source = JSON.parse(source);
            this.max_temp = source["max_temp"];
            this.min_temp = source["min_temp"];
            this.current_temp = source["current_temp"];
            this.current_stat = source["current_stat"];
            this.rain_day = source["rain_day"];
            this.rain_night = source["rain_night"];
            this.tmrw_rain_day = source["tmrw_rain_day"];
            this.tmrw_rain_night = source["tmrw_rain_night"];
            this.fine_dust = source["fine_dust"];
            this.ultra_dust = source["ultra_dust"];
            this.uv = source["uv"];
            this.sunset = source["sunset"];
            this.icon = source["icon"];
        }
    }
}