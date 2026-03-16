import './App.css'
import {ThemeProvider} from "@/providers/theme-provider.tsx";
import {ModeToggle} from "@/components/mode-toggle";
import Layout from "@/app/layout.tsx";

function App() {

  return (
    <>
        <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
            <Layout>
                <ModeToggle/>
            </Layout>
        </ThemeProvider>
    </>
  )
}

export default App
