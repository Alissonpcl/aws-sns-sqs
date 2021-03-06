/*
 * This file was generated by the Gradle 'init' task.
 *
 * This generated file contains a sample Java project to get you started.
 * For more details take a look at the Java Quickstart chapter in the Gradle
 * User Manual available at https://docs.gradle.org/6.2.2/userguide/tutorial_java_projects.html
 */

plugins {
    // Apply the java plugin to add support for Java
    java

    // Apply the application plugin to add support for building a CLI application.
    application
}

repositories {
    // Use jcenter for resolving dependencies.
    // You can declare any Maven/Ivy/file repository here.
    jcenter()
}

dependencies {

    implementation("com.google.code.gson:gson:2.8.6")

    // This dependency is used by the application.
    implementation("com.google.guava:guava:28.1-jre")

    // Use JUnit Jupiter API for testing.
    testImplementation("org.junit.jupiter:junit-jupiter-api:5.5.2")

    // Use JUnit Jupiter Engine for testing.
    testRuntimeOnly("org.junit.jupiter:junit-jupiter-engine:5.5.2")

    // Dependencias para AWS SDK, SNS e SQS
    implementation(platform("software.amazon.awssdk:bom:2.14.21"))
    implementation ("software.amazon.awssdk:sns")
    implementation ("software.amazon.awssdk:sqs")
}

application {
    // Define the main class for the application.
    mainClassName = "br.com.alissonlima.App"
}

val test by tasks.getting(Test::class) {
    // Use junit platform for unit tests
    useJUnitPlatform()
}
