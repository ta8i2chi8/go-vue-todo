<template>
    <div id="todos">
        <h1>TODO</h1>
        <div class='todo-list'>
            <table>
                <tr>
                    <th>TODO</th>
                    <th>CONTEXT</th>
                    <th>LIMIT</th>
                    <th>UPDATED_AT</th>
                    <th>DELETE</th>
                </tr>
                <tr v-for="todo in todos" :key=todo.id>
                    <td>{{ todo.title }}</td>
                    <td>{{ todo.context }}</td>
                    <td>{{ todo.limit_date }}</td>
                    <td>{{ todo.updated_at }}</td>
                    <td id="edit" @click="updateTodo(todo.id)">
                        <button>編集</button>
                    </td>
                    <td id="delete" @click="deleteTodo(todo.id)">
                        <button>削除</button>
                    </td>
                </tr>
            </table>
        </div>

        <input type="text" placeholder="title" v-model="form.title">
        <input type="text" placeholder="context" v-model="form.context">
        <input type="datetime-local" placeholder="datetime" @input="updateTime">
        <button @click="postTodo">post</button>

        <p>form</p>
        <p>{{ form }}</p>

        <p>todos</p>
        <p>{{ todos }}</p>
    </div> 
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { Todo } from '../types/todo'
import { APITodoRepository } from '../infra/todo_infra'

type DataType = {
    form: {
        title: string;
        context: string;
        limit_date: string;
    };
    todos: Todo[];
    message: string
}

interface HTMLInputEvent extends Event {
    target: HTMLInputElement;
}

export default defineComponent({
    data(): DataType {
        return {
            form: {
                title: '',     
                context: '',
                limit_date: '',
            },
            todos: [],
            message: ''
        }
    },
    mounted() {
        this.getTodos();
    },
    methods: {
        async getTodos(): Promise<void> {
            this.todos =  await this.$repos.TodoRepository.Get();
        },
        async createTodo(): Promise<void> {
            if (this.form.title !== '' && this.form.context !== '' && this.form.limit_date !== '') {
                this.message = await this.$repos.TodoRepository.Create(this.form);                
                this.todos =  await this.$repos.TodoRepository.Get();
                this.form = {
                    title: '',
                    context: '',
                    limit_date: '',
                };
            } else {
                this.message = 'Error！　必ず全て入力してください！';
            }
        },
        async updateTodo(id: number): Promise<void> {
            if (this.form.title !== '' && this.form.context !== '' && this.form.limit_date !== '') {
                this.message = await this.$repos.TodoRepository.Update(id, this.form);                
                this.todos =  await this.$repos.TodoRepository.Get();
                this.form = {
                    title: '',
                    context: '',
                    limit_date: '',
                };
            } else {
                this.message = 'Error！　必ず全て入力してください！';
            }
        },
        async deleteTodo(id: number): Promise<void> {
            this.message = await this.$repos.TodoRepository.Delete(id);
        },
        updateTime(event: HTMLInputEvent) {
            this.form.limit_date = event.target.value.replace(/T/, ' ').replace(/\//, '-') + ':00';
        },
    }
});
</script>

<style lang="scss" scoped>

</style>