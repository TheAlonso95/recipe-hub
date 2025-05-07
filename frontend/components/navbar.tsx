"use client";

import Link from "next/link";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";
import { auth } from "@/lib/auth";

export default function Navbar() {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [isLoading, setIsLoading] = useState(true);
  const router = useRouter();

  useEffect(() => {
    // Check authentication status on component mount
    setIsAuthenticated(auth.isAuthenticated());
    setIsLoading(false);
  }, []);

}, []);

  const handleLogout = async () => {
    try {
      await auth.logout();
      setIsAuthenticated(false);
      router.refresh();
    } catch (error) {
      console.error("Logout failed:", error);
      // TODO: Implement error handling for failed logout
    }
  };

useEffect(() => {
    // Check authentication status on component mount
    setIsAuthenticated(auth.isAuthenticated());
    setIsLoading(false);
  }, []);

  const handleLogout = async () => {
    try {
      await auth.logout();
      setIsAuthenticated(false);
      router.refresh();
    } catch (error) {
      console.error("Logout failed:", error);
      // TODO: Implement error handling for failed logout
    }
  };

  return (
    <nav className="bg-white shadow-sm border-b">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between h-16">
    await auth.logout();
    setIsAuthenticated(false);
    router.refresh();
  };

  return (
    <nav className="bg-white shadow-sm border-b">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between h-16">
          <div className="flex">
            <div className="flex-shrink-0 flex items-center">
              <Link href="/" className="font-bold text-xl">
                NextJS App
              </Link>
            </div>
          </div>
          <div className="flex items-center">
            {isLoading ? (
              <div>Loading...</div>
            ) : isAuthenticated ? (
              <div className="flex items-center space-x-4">
                <Button 
                  onClick={handleLogout} 
                  variant="outline"
                >
                  Logout
                </Button>
              </div>
            ) : (
              <div className="flex items-center space-x-4">
                <Link href="/login">
                  <Button variant="outline">Login</Button>
                </Link>
                <Link href="/register">
                  <Button>Register</Button>
                </Link>
              </div>
            )}
          </div>
        </div>
      </div>
    </nav>
  );
}