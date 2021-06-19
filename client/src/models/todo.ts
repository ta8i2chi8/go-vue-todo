class Todo {
    id: number;
    title: string;
    context: string;
    limit_date: string;
    created_at: string;
    updated_at: string;

    constructor(id: number, title: string, context: string, limit_date: string, create_at: string, updated_at: string) {
        this.id = id;
        this.title = title;
        this.context = context;
        this.limit_date = limit_date;
        this.created_at = create_at;
        this.updated_at = updated_at;
    }

}