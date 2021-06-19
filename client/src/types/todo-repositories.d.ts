import { ComponentCustomProperties } from 'vue'
import TodoRepository from '@/repository/todo_repository'

type Repositories = {
    TodoRepository: TodoRepository
};

declare module '@vue/runtime-core' {
    interface ComponentCustomProperties {
        $repos: Repositories;
    }
};