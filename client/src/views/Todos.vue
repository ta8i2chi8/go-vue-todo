<template>
    <div id="todos">
        <h1>TODO</h1>
        <article class='todo-list'>
            <table>
                <tr>
                    <th>TODO</th>
                    <th>CONTEXT</th>
                    <th>LIMIT</th>
                    <th>UPDATED_AT</th>
                    <th>STATUS</th>
                </tr>
                <tr v-for="todo in todos" :key=todo.id>
                    <td>{{ todo.title }}</td>
                    <td>{{ todo.context }}</td>
                    <td>{{ todo.limit_date }}</td>
                    <td>{{ todo.updated_at }}</td>
                    <td id="delete" @click="deleteTodo(todo.id)">delete</td>
                </tr>
            </table>
        </article>

        <input type="text" placeholder="title" v-model="form.title">
        <input type="text" placeholder="context" v-model="form.context">
        <input type="datetime-local" placeholder="datetime" @input="updateValue">
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
// import { APITodoRepository } from '@/infra/todo_infra'

type State = {
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
    data(): State {
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
    // created: {
    //     Get()
    // },
    methods: {
        updateValue(event: HTMLInputEvent) {
            this.form.limit_date = event.target.value.replace(/T/, ' ').replace(/\//, '-') + ':00';
        },
    }
});
</script>

<style lang="scss" scoped>
</style>