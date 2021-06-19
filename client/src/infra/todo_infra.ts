import axios from 'axios'
import { Todo, Form } from '@/types/todo'
import TodoRepository from '@/repository/todo_repository';

export class APITodoRepository implements TodoRepository {
    async Get() {
        const res = await axios.get<Todo[]>('http://localhost:8888/todos');
        const todos = res.data;
        return todos;
    }
    async Create(form: Form) {
        const res = await axios.post<{message: string}>('http://localhost:8888/todos', form);
        const message = res.data.message;
        return message;
    }
    async Update(id: number, form: Form) {
        const url = 'http://localhost:8888/todos/' + id;

        const res = await axios.put<{message: string}>(url, form);
        const message = res.data.message;
        return message;
    }
    async Delete(id: number) {
        const url = 'http://localhost:8888/todos/' + id;

        const res = await axios.delete<{message: string}>(url);
        const message = res.data.message;
        return message;
    }
}