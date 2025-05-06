import { NextResponse } from 'next/server'
import type { NextRequest } from 'next/server'

// This middleware runs on the edge and can protect routes
// or redirect users based on authentication status
export function middleware(request: NextRequest) {
  const token = request.cookies.get('token')?.value || request.headers.get('authorization')?.split(' ')[1]
  
  // Path that requires authentication
  const isProtectedRoute = request.nextUrl.pathname.startsWith('/protected')
  
  // Auth-related paths
  const isAuthRoute = request.nextUrl.pathname.startsWith('/login') || 
                      request.nextUrl.pathname.startsWith('/register')

  // If trying to access protected route without authentication, redirect to login
  if (isProtectedRoute && !token) {
    return NextResponse.redirect(new URL('/login', request.url))
  }

  // If already authenticated and trying to access login/register, redirect to home
  if (isAuthRoute && token) {
    return NextResponse.redirect(new URL('/', request.url))
  }

  return NextResponse.next()
}

// Configure which paths should trigger this middleware
export const config = {
  matcher: ['/protected/:path*', '/login', '/register'],
}