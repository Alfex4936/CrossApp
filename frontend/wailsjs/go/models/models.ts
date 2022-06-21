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
}