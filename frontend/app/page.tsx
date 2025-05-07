import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import Link from "next/link"

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center pt-16 px-4">
      <div className="w-full max-w-5xl">
        <h1 className="text-4xl font-bold text-center mb-8">
          Welcome to Next.js with Authentication
        </h1>
        
        <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-2">
          <Card className="shadow-md">
            <CardHeader>
              <CardTitle>Login to Your Account</CardTitle>
              <CardDescription>
                Access your account with your credentials
              </CardDescription>
            </CardHeader>
            <CardContent>
              <p className="mb-4">
                If you already have an account, you can log in to access your dashboard and personalized features.
              </p>
              <Link href="/login">
                <Button className="w-full">Go to Login</Button>
              </Link>
            </CardContent>
          </Card>

          <Card className="shadow-md">
            <CardHeader>
              <CardTitle>Create a New Account</CardTitle>
              <CardDescription>
                Register to get started with our services
              </CardDescription>
            </CardHeader>
            <CardContent>
              <p className="mb-4">
                New to our platform? Create an account to unlock all features and start your journey.
              </p>
              <Link href="/register">
                <Button className="w-full" variant="outline">Register Now</Button>
              </Link>
            </CardContent>
          </Card>
        </div>

        <div className="mt-10 space-y-4">
          <h2 className="text-2xl font-semibold">Authentication Features</h2>
          <ul className="list-disc pl-6 space-y-2">
            <li>Secure user registration with email validation</li>
            <li>Login with JWT authentication</li>
            <li>Protected API routes for authenticated users</li>
            <li>Client-side form validation using Zod</li>
            <li>Modern UI components with shadcn/ui and Tailwind CSS</li>
          </ul>
        </div>
      </div>
    </main>
  )
}