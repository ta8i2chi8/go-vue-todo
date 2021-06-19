import { createApp } from 'vue'
import App from './App.vue'
import { APITodoRepository } from './infra/todo_infra';
import router from './router'

const app = createApp(App).use(router);

app.config.globalProperties.$repos = {
    TodoRepository: new APITodoRepository
};

app.mount('#app');