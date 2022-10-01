import { Form } from '@/types/todo'
import Todo from '@/models/todo'

export default interface TodoRepository {
    Get(): Promise<Todo[]>;
    Create(form: Form): Promise<string>;
    Update(id: number, form: Form): Promise<string>;
    Delete(id: number): Promise<string>;
}