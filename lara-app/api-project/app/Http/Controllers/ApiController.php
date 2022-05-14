<?php

namespace App\Http\Controllers;

use App\Models\Student;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;

class ApiController extends Controller
{

    public function getAllStudents(): \Illuminate\Http\Response|\Illuminate\Contracts\Foundation\Application|\Illuminate\Contracts\Routing\ResponseFactory
    {
        // logic to get all students goes here
        $students = Student::get()->toJson(JSON_PRETTY_PRINT);
        return response($students, 200);
    }

    public function createStudent(Request $request): \Illuminate\Http\JsonResponse
    {
        // logic to create a student record goes here
        $student = new Student();
        $student->name = $request->name;
        $student->course = $request->course;
        $student->save();

        return response()->json([
            "message" => "student record created"
        ], 201);
    }

    public function getStudent($id)
    {
        // logic to get a student record goes here
    }

    public function updateStudent(Request $request, $id) {
        // logic to update a student recotd goes here
    }

    public function deleteStudent($id) {
        // logic to delete a student recotd goes here
    }
}
