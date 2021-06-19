import { Todo, Form } from '@/types/todo'

export default interface TodoRepository {
    Get(): Promise<Todo[]>;
    Create(form: Form): Promise<string>;
    Update(id: number, form: Form): Promise<string>;
    Delete(id: number): Promise<string>;
}