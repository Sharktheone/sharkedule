import {logo} from "@/utils/logo"
import {Input} from "@mantine/core"
import styles from "./styles.module.scss"


export default function Login() {
    return (
        <div>
            <div className={styles.loginContainer}>
                <div className={styles.logoContainer}>
                    <img className={styles.logo} src={logo()} alt="logo"/>
                    <h1>Sharkedule</h1>
                </div>
                <h1>Sign in</h1>
                <form>
                    <Input type="text" placeholder="Username"/>
                    <Input type="password" placeholder="Password"/>
                    <button type="submit">Sign in</button>
                </form>
            </div>
        </div>
    )
}