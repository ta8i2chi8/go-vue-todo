import axios from 'axios'
import { Todo, Form } from '@/types/todo'
import TodoRepository from '@/repository/todo_repository';

export class APITodoRepository implements TodoRepository {
    async Get() {
        try {
            const res = await axios.get<Todo[]>('http://localhost:8888/todos');
            const todos = res.data
            return todos
        } 
        catch (err: any) {
            if (err.response !== undefined) {
                return err.response.data.error;
            }            
        }  
    }
    async Create(form: Form) {
        try {
            const res = await axios.post<{message: string}>('http://localhost:8888/todos', form)
            const message = res.data.message
            return message
        }
        catch (err: any) {
            if (err.response !== undefined) {
                return err.response.data.error;
            }            
        }
    }
    async Update(id: number, form: Form) {
        const url = 'http://localhost:8888/todos/' + id;

        try {
            const res = await axios.put<{message: string}>(url, form)
            const message = res.data.message
            return message
        }
        catch (err: any) {
            if (err.response !== undefined) {
                return err.response.data.error;
            }            
        }
    }
    async Delete(id: number) {
        const url = 'http://localhost:8888/todos/' + id;

        try {
            const res = await axios.delete<{message: string}>(url)
            const message = res.data.message
            return message
        }
        catch (err: any) {
            if (err.response !== undefined) {
                return err.response.data.error;
            }            
        }
    }
}