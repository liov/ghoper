#include "test_jni_JNI.h"

//gcc test_jni_JNI.c -shared -o /f/tool/dll/hello.dll -I /e/jdk-12/include -I /e/jdk-12/include/win32

JNIEXPORT void JNICALL Java_test_jni_JNI_testHelloVoid (JNIEnv *env, jobject obj) {
  puts("hello world return void");
}

JNIEXPORT jstring JNICALL Java_test_jni_JNI_testHello (JNIEnv *env, jobject obj) {
    const char *p ="hello world return jstring";
  return (*env)->NewStringUTF(env,p);
}
