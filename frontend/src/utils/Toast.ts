import {message} from 'ant-design-vue';


function ToastInfo(msg: string) {
    message.info(msg);
}

function ToastError(msg: string) {
    message.error(msg);
}


export {ToastError, ToastInfo}