import { ref } from 'vue'
import axios from 'axios'

export function useQrGenerator() {
    const base64Result = ref<string>('')

    async function generateQr(content: string, variant: string) {
        if (!content) {
            base64Result.value = ''
            return false
        }

        try {
            const response = await axios.post('/generate-qr', {
                content,
                variant
            })
            base64Result.value = response.data.base64
            return true
        } catch (error) {
            console.error(error)
            return false
        }
    }

    return {
        base64Result,
        generateQr
    }
}
