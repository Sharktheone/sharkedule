import {logo} from "@/utils/logo"
import {Input} from "@mantine/core"
import styles from "./styles.module.scss"
import {useState} from "react"
import {IconArrowBigRight} from "@tabler/icons-react"


export default function Login() {

    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")

    function login(e: React.FormEvent<HTMLFormElement>) {
        e.preventDefault()
        //TODO: login api, hash password with sha512
    }


    return (
        <div>
            <div className={styles.loginContainer}>
                <div className={styles.logoContainer}>
                    <img className={styles.logo} src={logo()} alt="logo"/>
                    <h1>Sharkedule</h1>
                </div>
                <h1>Sign in</h1>
                <form className={styles.loginForm} onSubmit={login}>
                    <div>
                        <label htmlFor="username">Username or Email</label>
                        <Input name="username" id="username" type="text" placeholder="Username"
                               onChange={e => setUsername(e.currentTarget.value)}
                               value={username}
                        />
                    </div>
                    <div>
                        <label htmlFor="password">Password</label>
                        <Input name="password" id="password" type="password" placeholder="Password"
                               onChange={e => setPassword(e.currentTarget.value)}
                               value={password}
                        />
                    </div>
                    <button className={styles.submit} type="submit">
                        Sign in
                        <IconArrowBigRight/>
                    </button>
                </form>
            </div>
        </div>
    )
}