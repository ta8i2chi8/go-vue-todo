<template>
    <div id="todos">
        <h1>TODOアプリ</h1>
        <div class='todo-list' v-if="!isEdit">
            <table>
                <tr>
                    <th class="title">TODO</th>
                    <th class="context">CONTEXT</th>
                    <th>LIMIT</th>
                    <th>UPDATED_AT</th>
                    <th>EDIT</th>
                    <th>DELETE</th>
                </tr>
                <tr v-for="todo in todos" :key=todo.id>
                    <td class="title">{{ todo.title }}</td>
                    <td class="context">{{ todo.context }}</td>
                    <td>{{ todo.limit_date }}</td>
                    <td>{{ todo.updated_at }}</td>
                    <td id="edit">
                        <button type="button" @click="moveToEdit(todo.id, todo.title, todo.context, todo.limit_date)">編集</button>
                    </td>
                    <td id="delete">
                        <button type="button" @click="deleteTodo(todo.id)">削除</button>
                    </td>
                </tr>
            </table>
        </div>

        <p>{{ message }}</p>

        <form v-if="!isEdit">
            <h2>タスク作成</h2>
            <input type="text" placeholder="title" v-model="form.title">
            <input type="text" placeholder="context" v-model="form.context">
            <input type="datetime-local" autocomplete="off" @input="updateTime">
            
            <button type="button" class="btn post" @click="createTodo">追加</button>
        </form>
        
        <form v-else>
            <h2>タスク編集</h2>
            <input type="text" placeholder="title" v-model="form.title">
            <input type="text" placeholder="context" v-model="form.context">
            <input type="datetime-local" :value="initLimitDate" autocomplete="off" @input="updateTime">
            
            <button type="button" class="btn edit" @click="updateTodo">編集</button>
            <button type="button" class="btn" @click="isEdit = !isEdit">戻る</button>
        </form>
    </div> 
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Todo from '@/models/todo';

type DataType = {
    form: {
        title: string;
        context: string;
        limit_date: string;
    };
    todos: Todo[];
    message: string;
    isEdit: boolean;
    editId: number;
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
            message: '',
            isEdit: false,
            editId: 0,
        }
    },
    computed: {
        initLimitDate(): string {
            return this.form.limit_date.replace(' ', 'T');           
        }
    },
    created() {
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
                this.clear();
            } else {
                this.message = 'Error！　必ず全て入力してください！';
            }
        },
        async updateTodo(): Promise<void> {
            if (this.form.title !== '' && this.form.context !== '' && this.form.limit_date !== '') {
                this.message = await this.$repos.TodoRepository.Update(this.editId, this.form);                
                this.todos =  await this.$repos.TodoRepository.Get();
                this.isEdit = !this.isEdit;
            } else {
                this.message = 'Error！　必ず全て入力してください！';
            }
        },
        async deleteTodo(id: number): Promise<void> {
            this.message = await this.$repos.TodoRepository.Delete(id);
            this.todos =  await this.$repos.TodoRepository.Get();
        },
        updateTime(event: HTMLInputEvent) {
            this.form.limit_date = event.target.value.replace(/T/, ' ').replace(/\//, '-') + ':00';
        },
        moveToEdit(id: number, title: string, context: string, limit_date: string) {
            this.isEdit = !this.isEdit;
            this.editId = id;
            this.form = {
                title: title,
                context: context,
                limit_date: limit_date,
            };
        },
        clear() {
            this.form = {
                title: '',
                context: '',
                limit_date: '',
            };
            this.editId = 0;
        },
    }
});
</script>

<style lang="sass" scoped>
.todo-list
    table
        width: 65%
        margin: 0 auto

        .title
            width: 15%

        .context
            width: 25%
        
        th:first-child
            border-radius: 5px 0 0 0

        th:last-child
            border-radius: 0 5px 0 0
            border-right: 1px solid #3c6690

        th
            text-align: center
            color: white
            background: linear-gradient(#829ebc,#225588)
            border-left: 1px solid #3c6690
            border-top: 1px solid #3c6690
            border-bottom: 1px solid #3c6690
            box-shadow: 0px 1px 1px rgba(255,255,255,0.3) inset
            padding: 10px 0

        td
            text-align: center
            border-left: 1px solid #a8b7c5
            border-bottom: 1px solid #a8b7c5
            border-top: none
            box-shadow: 0px -3px 5px 1px #eee inset
            width: 25%
            padding: 10px 0


        td:last-child
            border-right: 1px solid #a8b7c5

form
    width: 40%
    margin: 50px auto
    min-width: 100px

    input
        float: left
        width: 100%
        max-width: 100%
        border: none
        margin: 0.5rem 0
        padding: 0.5rem 1rem
        border-radius: 0.3rem
        background: darken(#f9f9f9, 10%)
        color: darken(#f9f9f9, 50%)

        &::placeholder 
            color: darken(#f9f9f9, 50%)

    .post
        display: block
        margin: 30px auto
    
    .edit
        margin-right: 10px
</style>